/*
Authentication domain will be set here
*/
package domain

import (
	"github.com/gin-gonic/gin"
	"github.com/louistwiice/go/basicwithent/entity"
)

type AuthService interface {
	Create(u *entity.UserCreateUpdate) error
	GetByID(id string) (*entity.UserDisplay, string, error)
	SearchUser(identifier string) (*entity.UserDisplay, string, error)
	UpdateAuthenticationDate(u *entity.UserDisplay) error
}

type AuthController interface {
	register(ctx *gin.Context)
	login(ctx *gin.Context)
	refreshToken(ctx *gin.Context)
	logout(ctx *gin.Context)
}