package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	entsql "entgo.io/ent/dialect/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"github.com/louistwiice/go/basicwithent/api/controllers"
	"github.com/louistwiice/go/basicwithent/api/middlewares"
	"github.com/louistwiice/go/basicwithent/configs"
	"github.com/louistwiice/go/basicwithent/ent"
	"github.com/louistwiice/go/basicwithent/ent/migrate"
	"github.com/louistwiice/go/basicwithent/repository"
	"github.com/louistwiice/go/basicwithent/usecase/authentication"
	"github.com/louistwiice/go/basicwithent/usecase/user"
)

// To load .env file
func init() {
	configs.Initialize()
}

// To create new database connection
func open(source string) (*ent.Client, error) {
    db, err := sql.Open("mysql", source)
    if err != nil {
        return nil, err
    }
    db.SetMaxIdleConns(10)
    db.SetMaxOpenConns(100)
    db.SetConnMaxLifetime(time.Hour)
    // Create an ent.Driver from `db`.
    drv := entsql.OpenDB("mysql", db)
    return ent.NewClient(ent.Driver(drv)), nil
}

func main() {
	log.Println("Server starting ...")

	// Start by connecting to database
	dbSource := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", configs.GetString("DB_USER"), configs.GetString("DB_PASSWORD"), configs.GetString("DB_HOST"), configs.GetString("DB_NAME"))
	db, err := open(dbSource)
	if err!= nil {
		log.Panic("Database error: ... ", err.Error())
	}
	defer db.Close()

	// Run the automatic migration tool to create all schema resources.
	ctx := context.Background()
	err = db.Schema.Create(
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

	userController := controllers.NewUserController(userService)
	authController := controllers.NewAuthController(authService)
	middlwareController := middlewares.NewMiddlewareControllers()
	
	app := gin.Default()

	api_v1 := app.Group("api/v1")
	authController.MakeAuthHandlers(api_v1.Group("auth/"))
	userController.MakeUserHandlers(api_v1.Group("user/"))

	api_auth := app.Group("api/v1/in")
	api_auth.Use(middlwareController.JwAuthtMiddleware())
	userController.MakeUserHandlers(api_auth.Group("user/"))

	app.Run(configs.GetString("SERVER_PORT"))
}

