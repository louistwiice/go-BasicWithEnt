package handler_users

import (
	"github.com/louistwiice/go/basicwithent/entity"
	"github.com/louistwiice/go/basicwithent/repository"
	"github.com/louistwiice/go/basicwithent/usecase/authentication"
	"github.com/louistwiice/go/basicwithent/usecase/user"
)

func NewUserRouters(server *entity.Routers) {
	userRepo := repository.NewUserClient(server.Database)

	userService := user.NewUserService(userRepo)
	authService := authentication.NewAuthService(userRepo)
	
	userController := NewUserController(userService)
	authController := NewAuthController(authService)

	// Basic authentication system
	api_connection := server.OpenApp.Group("auth")
	api_connection.POST("login", authController.login)
	api_connection.POST("register", authController.register)
	api_connection.POST("refresh", authController.refreshToken)
	api_connection.GET("logout", authController.logout)

	// User management
	api_user := server.OpenApp.Group("user")
	api_user.GET("", userController.listUsers)
	api_user.GET(":id", userController.getUser)
	api_user.PUT(":id", userController.updateUser)
	api_user.POST(":id/reset_password", userController.updatePassword)
	api_user.DELETE(":id", userController.deleteUser)

	// Authentication required
	api_auth := server.RestrictedApp.Group("user")
	api_auth.GET("", userController.listUsers)
	api_auth.DELETE(":id", userController.deleteUser)
}