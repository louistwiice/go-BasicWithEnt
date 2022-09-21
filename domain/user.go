package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/louistwiice/go/basicwithent/entity"
)

type UserRepository interface {
	List() ([]*entity.UserDisplay, error)
	Create(u *entity.UserCreateUpdate) error
	GetByID(id string) (*entity.UserDisplay, string, error)
	SearchUser(identifier string) (*entity.UserDisplay, string, error)
	UpdateInfo(u *entity.UserCreateUpdate) error
	UpdatePassword(u *entity.UserCreateUpdate) error
	UpdateAuthenticationDate(u *entity.UserDisplay) error
	Delete(id string) error
}

type UserService interface {
	List() ([]*entity.UserDisplay, error)
	Create(u *entity.UserCreateUpdate) error
	GetByID(id string) (*entity.UserDisplay, string, error)
	SearchUser(identifier string) (*entity.UserDisplay, string, error)
	UpdateUser(u *entity.UserCreateUpdate) error
	UpdatePassword(u *entity.UserCreateUpdate) error
	Delete(id string) error
}

type UserController interface {
	listUser(ctx *gin.Context)
	getUser(ctx *gin.Context)
	updateUser(ctx *gin.Context)
	updatePassword(ctx *gin.Context)
	deleteUser(ctx *gin.Context)
}