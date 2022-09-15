package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/louistwiice/go/basicwithent/entity"
)

type UserRepository interface {
	List() ([]*entity.UserDisplay, error)
	Create(u *entity.UserCreateUpdate) error
	Get(id string) (*entity.UserDisplay, string, error)
	UpdateInfo(u *entity.UserCreateUpdate) error
	UpdatePassword(u *entity.UserCreateUpdate) error
	//Delete(id int) error
}

type UserService interface {
	List() ([]*entity.UserDisplay, error)
	Create(u *entity.UserCreateUpdate) error
	Get(id string) (*entity.UserDisplay, string, error)
	UpdateUser(u *entity.UserCreateUpdate) error
	UpdatePassword(u *entity.UserCreateUpdate) error
	//Delete(id int) error
}

type UserController interface {
	ListUser(ctx *gin.Context)
	CreateUser(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
	UpdatePassword(ctx *gin.Context)
}
