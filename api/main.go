package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	logger "github.com/rs/zerolog/log"

	"github.com/louistwiice/go/basicwithent/api/handlers/users"
	"github.com/louistwiice/go/basicwithent/api/middlewares"
	"github.com/louistwiice/go/basicwithent/configs"
	"github.com/louistwiice/go/basicwithent/ent/migrate"
	"github.com/louistwiice/go/basicwithent/repository"
	"github.com/louistwiice/go/basicwithent/usecase/authentication"
	"github.com/louistwiice/go/basicwithent/usecase/user"
)

// To load .env file
func init() {
	configs.Initialize()
}

func main() {
	logger.Info().Msg("Server starting ...")
	conf := configs.LoadConfigEnv()

	// Start by connecting to database
	db := configs.NewDBConnection()
	defer db.Close()

	// Run the automatic migration tool to create all schema resources.
	ctx := context.Background()
	err := db.Schema.Create(
        ctx, 
        migrate.WithDropIndex(true),
        migrate.WithDropColumn(true), 
    )
	if err!= nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	userRepo := repository.NewUserClient(db)

	userService := user.NewUserService(userRepo)
	authService := authentication.NewAuthService(userRepo)

	userController := handler_users.NewUserController(userService)
	authController := handler_users.NewAuthController(authService)
	middlwareController := middlewares.NewMiddlewareControllers(authService)
	
	app := gin.Default()

	api_v1 := app.Group("api/v1")
	authController.MakeAuthHandlers(api_v1.Group("auth/"))
	userController.MakeUserHandlers(api_v1.Group("user/"))

	api_auth := app.Group("api/v1/in")
	api_auth.Use(middlwareController.JwAuthtMiddleware())
	userController.MakeUserHandlers(api_auth.Group("user/"))

	logger.Info().Msg("Server ready to go ...")
	app.Run(conf.ServerPort)
}

