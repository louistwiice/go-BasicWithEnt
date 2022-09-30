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
	"github.com/louistwiice/go/basicwithent/entity"
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

	app := gin.Default()
	api_v1 := app.Group("api/v1")
	api_restricted := app.Group("api/v1/in")

	router_base := &entity.RouterBase{
		Database: db,
		OpenApp: api_v1,
	}
	router := &entity.Routers{
		RouterBase: *router_base,
		RestrictedApp: api_restricted,
	}

	middlewareController := middlewares.NewMiddlewareRouters(router)
	api_restricted.Use(middlewareController.JwAuthtMiddleware())

	handler_users.NewUserRouters(router)

	logger.Info().Msg("Server ready to go ...")
	app.Run(conf.ServerPort)
}

