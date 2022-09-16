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
}

type UserService interface {
	List() ([]*entity.UserDisplay, error)
	Create(u *entity.UserCreateUpdate) error
	GetByID(id string) (*entity.UserDisplay, string, error)
	SearchUser(identifier string) (*entity.UserDisplay, string, error)
	UpdateUser(u *entity.UserCreateUpdate) error
	UpdatePassword(u *entity.UserCreateUpdate) error
}
/*
type AuthService interface {
	Create(u *entity.UserCreateUpdate) error
	GetByID(id string) (*entity.UserDisplay, string, error)
	SearchUser(identifier string) (*entity.UserDisplay, string, error)
}*/

type UserController interface {
	listUser(ctx *gin.Context)
	createUser(ctx *gin.Context)
	getUser(ctx *gin.Context)
	updateUser(ctx *gin.Context)
	updatePassword(ctx *gin.Context)
}

/*
type AuthController interface {
	register(ctx *gin.Context)
	login(ctx *gin.Context)
	refreshToken(ctx *gin.Context)
}*/