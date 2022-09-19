package user

import (
	"github.com/louistwiice/go/basicwithent/entity"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct{
	mock.Mock
}

func (m *MockUserService) List() ([]*entity.UserDisplay, error) {
	args := m.Called()
	
	r0 := args.Get(0).([]*entity.UserDisplay)
	r1 := args.Get(1).(error)

	return r0, r1
}

func (m *MockUserService) Create(u *entity.UserCreateUpdate) error {
	args := m.Called(u)

	var r0 error
	if rf, ok := args.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = args.Error(0)
	}

	return r0
}

func (m *MockUserService) GetByID(id string) (*entity.UserDisplay, string, error) {
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

func (m *MockUserService) UpdateUser(u *entity.UserCreateUpdate) error {
	args := m.Called(u)

	var r0 error
	if rf, ok := args.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = args.Error(0)
	}

	return r0
}

func (m *MockUserService) UpdatePassword(u *entity.UserCreateUpdate) error {
	args := m.Called(u)

	var r0 error
	if rf, ok := args.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = args.Error(0)
	}

	return r0
}

func (m *MockUserService) SearchUser(identifier string) (*entity.UserDisplay, string, error) {
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

	var r2 error
	if rf, ok := args.Get(2).(func() error); ok {
		r2 = rf()
	} else {
		r2 = args.Error(2)
	}

	return r0, r1, r2
}