package repository

import (
	"testing"

	"github.com/louistwiice/go/basicwithent/ent"
	"github.com/louistwiice/go/basicwithent/ent/enttest"
	"github.com/louistwiice/go/basicwithent/ent/migrate"
	"github.com/louistwiice/go/basicwithent/entity"
	"github.com/louistwiice/go/basicwithent/mocks"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func Test_Create(t *testing.T) {
	opts := []enttest.Option{
        enttest.WithOptions(ent.Log(t.Log)),
        enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
    }
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
	defer client.Close()

	repo := NewUserClient(client)

	// Test for Create user
	userCreate := mocks.GenerateFixture().UserCreate1

	response_create := repo.Create(userCreate)
	assert.Nil(t, response_create)
}

func Test_Get(t *testing.T) {
	opts := []enttest.Option{
        enttest.WithOptions(ent.Log(t.Log)),
        enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
    }
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
	defer client.Close()

	repo := NewUserClient(client)
	email := "mike@mail.com"
	username := "mike"

	userCreate := mocks.GenerateFixture().UserCreate1
	_ = repo.Create(userCreate)

	// Test for GetByID
	response_getbyID, _, err := repo.GetByID(userCreate.ID)
	assert.Nil(t, err)
	assert.Equal(t, userCreate.ID, response_getbyID.ID)
	assert.Equal(t, userCreate.Email, response_getbyID.Email)
	assert.Equal(t, userCreate.LastName, response_getbyID.LastName)
	assert.Equal(t, userCreate.Username, response_getbyID.Username)

	_, _, err = repo.GetByID("483ed845-387f-4c50-9a64-fef701b4dbb8") // This ID does not exist in database
	assert.NotNil(t, err)
	assert.Equal(t, entity.ErrNotFound, err, "The ID does not exist. It should return not found")

	// Test for SearchUser
	response_search, _, err := repo.SearchUser(email)
	assert.Nil(t, err)
	assert.Equal(t, userCreate.ID, response_search.ID)
	assert.Equal(t, userCreate.Email, response_search.Email)
	assert.Equal(t, userCreate.LastName, response_search.LastName)
	assert.Equal(t, userCreate.Username, response_search.Username)

	_, _, err = repo.SearchUser("Steven")
	assert.Equal(t, entity.ErrNotFound, err, "The username does not exist. It should return not found")

	response_search, _, err = repo.SearchUser(username)
	assert.Nil(t, err)
	assert.Equal(t, userCreate.ID, response_search.ID)
	assert.Equal(t, userCreate.Email, response_search.Email)
	assert.Equal(t, userCreate.LastName, response_search.LastName)
	assert.Equal(t, userCreate.Username, response_search.Username)

	_, _, err = repo.SearchUser("steven@mail.com") 
	assert.Equal(t, entity.ErrNotFound, err, "The email does not exist. So it should be not found")
}

func Test_List(t *testing.T) {
	opts := []enttest.Option{
        enttest.WithOptions(ent.Log(t.Log)),
        enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
    }
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
	defer client.Close()

	repo := NewUserClient(client)

	userCreate := mocks.GenerateFixture().UserCreate2

	response_list, err := repo.List()
	assert.Nil(t, err)
	assert.Equal(t, 0, len(response_list))

	_ = repo.Create(userCreate)

	response_list, err = repo.List()
	assert.Nil(t, err)
	assert.Equal(t, 1, len(response_list))
	assert.Equal(t, userCreate.ID, response_list[0].ID)
	assert.Equal(t, userCreate.Email, response_list[0].Email)
}

func Test_UpdateInfo(t *testing.T) {
	opts := []enttest.Option{
        enttest.WithOptions(ent.Log(t.Log)),
        enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
    }
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
	defer client.Close()

	repo := NewUserClient(client)

	userCreate := mocks.GenerateFixture().UserCreate2

	err := repo.UpdateInfo(userCreate)
	assert.NotNil(t, err, "It should return an error as the DB is empty")

	_ = repo.Create(userCreate)

	// Test for List
	err = repo.UpdateInfo(userCreate)
	assert.Nil(t, err)
}

func Test_UpdatePassword(t *testing.T) {
	opts := []enttest.Option{
        enttest.WithOptions(ent.Log(t.Log)),
        enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
    }
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1", opts...)
	defer client.Close()

	repo := NewUserClient(client)
	new_password := "john_new_password"

	userCreate := mocks.GenerateFixture().UserCreate2

	err := repo.UpdatePassword(userCreate)
	assert.NotNil(t, err, "It should return an error as the DB is empty")

	_ = repo.Create(userCreate)

	userCreate.Password = new_password
	err = repo.UpdatePassword(userCreate)
	assert.Nil(t, err)
}