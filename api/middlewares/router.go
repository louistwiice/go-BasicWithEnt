package middlewares

import (
	"github.com/louistwiice/go/basicwithent/entity"
	"github.com/louistwiice/go/basicwithent/repository/user"
	"github.com/louistwiice/go/basicwithent/usecase/authentication"
)


func NewMiddlewareRouters(server *entity.Routers) *controller {
	userRepo := repository_user.NewUserClient(server.Database)
	authService := authentication.NewAuthService(userRepo)

	return NewMiddlewareControllers(authService)
}