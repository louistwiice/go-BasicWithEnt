package service_user

import (
	"testing"

	"github.com/louistwiice/go/basicwithent/entity"
	"github.com/louistwiice/go/basicwithent/mocks"
	"github.com/louistwiice/go/basicwithent/mocks/user"
	"github.com/stretchr/testify/assert"
)

func Test_List(t *testing.T) {
	u := mocks.GenerateFixture().UserList
	repo := user.MockUserRepo{}
	repo.On("List").Return(u, nil)
	
	service := NewUserService(&repo)
	response, err := service.List()

	assert.Nil(t, err)
	assert.Equal(t, 2, len(response))
	assert.Equal(t, u, response)
}

func Test_Create(t *testing.T) {
	u := mocks.GenerateFixture().UserCreate1
	repo := user.MockUserRepo{}
	repo.On("Create", u).Return(nil)
	
	service := NewUserService(&repo)
	err := service.Create(u)
	assert.Nil(t, err)
}

func Test_GetByID(t *testing.T) {

	t.Run("There should be no error when repo sent Nil", func(t *testing.T) {
		u := mocks.GenerateFixture().UserDisplay1
		u_password := mocks.GenerateFixture().User1Password
		repo := user.MockUserRepo{}
		repo.On("GetByID", u.ID).Return(u, u_password, nil).Once()
		
		service := NewUserService(&repo)
		response, password, err := service.GetByID(u.ID)
		assert.Nil(t, err)
		assert.Equal(t, u_password, password)
		assert.Equal(t, u, response)
	})

	t.Run("There should be an error when repo does not send a nil error", func(t *testing.T) {
		u := mocks.GenerateFixture().UserDisplay1
		repoNotNil := user.MockUserRepo{}
		repoNotNil.On("GetByID", u.ID).Return(&entity.UserDisplay{}, "", entity.ErrNotFound)
		
		service := NewUserService(&repoNotNil)
		response, password, err := service.GetByID(u.ID)
		assert.NotNil(t, err)
		assert.Equal(t, "", password)
		assert.Equal(t, &entity.UserDisplay{}, response)
		assert.Equal(t, entity.ErrNotFound, err)
	})
}

func Test_UpdateUser(t *testing.T) {
	u := mocks.GenerateFixture().UserCreate1
	repo := user.MockUserRepo{}
	repo.On("UpdateInfo", u).Return(nil)
	
	service := NewUserService(&repo)
	err := service.UpdateUser(u)
	assert.Nil(t, err)
}

func Test_UpdatePassword(t *testing.T) {
	u := mocks.GenerateFixture().UserCreate1
	repo := user.MockUserRepo{}
	repo.On("UpdatePassword", u).Return(nil)
	
	service := NewUserService(&repo)
	err := service.UpdatePassword(u)
	assert.Nil(t, err)
}

func Test_SearchUser(t *testing.T) {

	t.Run("can search a user if is valid", func(t *testing.T) {
		u := mocks.GenerateFixture().UserDisplay1
		u_password := mocks.GenerateFixture().User1Password
		repoNil := user.MockUserRepo{}
		repoNil.On("SearchUser", u.Email).Return(u, u_password, nil)

		service := NewUserService(&repoNil)
		response, password, err := service.SearchUser(u.Email)
		assert.Nil(t, err)
		assert.Equal(t, u_password, password)
		assert.Equal(t, u, response)
	})

	t.Run("can search a user if is valid", func(t *testing.T) {
		u := mocks.GenerateFixture().UserDisplay1
		repoNotNil := user.MockUserRepo{}
		repoNotNil.On("SearchUser", u.Email).Return(&entity.UserDisplay{}, "", entity.ErrNotFound)
	
		service := NewUserService(&repoNotNil)
		_, password, err := service.SearchUser(u.Email)
		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Equal(t, "", password)
	})		
}

func Test_Delete(t *testing.T) {

	t.Run("Return nil when repo returns nil", func(t *testing.T) {
		u := mocks.GenerateFixture().UserDisplay1
		repo := user.MockUserRepo{}
		repo.On("Delete", u.ID).Return(nil)

		service := NewUserService(&repo)
		err := service.Delete(u.ID)
		assert.Nil(t, err)
	})

	t.Run("Retrun an error when repo returns an erro", func(t *testing.T) {
		u := mocks.GenerateFixture().UserDisplay1
		repo := user.MockUserRepo{}
		repo.On("Delete", u.ID).Return(entity.ErrNotFound)
	
		service := NewUserService(&repo)
		err := service.Delete(u.ID)
		assert.NotNil(t, err)
	})		
}