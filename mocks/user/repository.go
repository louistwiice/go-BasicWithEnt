package user

import (
	"github.com/louistwiice/go/basicwithent/entity"
	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct{
	mock.Mock
}

func (m *MockUserRepo) List() ([]*entity.UserDisplay, error) {
	args := m.Called()

	var r0 []*entity.UserDisplay
	if rf, ok := args.Get(0).(func() []*entity.UserDisplay); ok {
		r0 = rf()
	} else {
		r0 = args.Get(0).([]*entity.UserDisplay)
	}

	r1 := args.Error(1)

	return r0, r1
}

func (m *MockUserRepo) Create(u *entity.UserCreateUpdate) error {
	args := m.Called(u)

	return args.Error(0)
}

func (m *MockUserRepo) GetByID(id string) (*entity.UserDisplay, string, error) {
	args := m.Called(id)

	var r0 *entity.UserDisplay
	if rf, ok := args.Get(0).(func() *entity.UserDisplay); ok {
		r0 = rf()
	} else {
		r0 = args.Get(0).(*entity.UserDisplay)
	}

	var r1 string
	if rf, ok := args.Get(1).(func() string); ok {
		r1 = rf()
	} else {
		r1 = args.Get(1).(string)
	}

	r2 := args.Error(2)

	return r0, r1, r2
}

func (m *MockUserRepo) UpdateInfo(u *entity.UserCreateUpdate) error {
	args := m.Called(u)

	return args.Error(0)
}

func (m *MockUserRepo) UpdatePassword(u *entity.UserCreateUpdate) error {
	args := m.Called(u)

	return args.Error(0)
}

func (m *MockUserRepo) UpdateAuthenticationDate(u *entity.UserDisplay) error {
	args := m.Called(u)

	return args.Error(0)
}

func (m *MockUserRepo) SearchUser(identifier string) (*entity.UserDisplay, string, error) {
	args := m.Called(identifier)

	var r0 *entity.UserDisplay
	if rf, ok := args.Get(0).(func() *entity.UserDisplay); ok {
		r0 = rf()
	} else {
		r0 = args.Get(0).(*entity.UserDisplay)
	}

	var r1 string
	if rf, ok := args.Get(1).(func() string); ok {
		r1 = rf()
	} else {
		r1 = args.Get(1).(string)
	}

	r2 := args.Error(2)

	return r0, r1, r2
}
