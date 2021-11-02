package main

import (
	env "github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"laughing-succostash/internal/core/database"
	"laughing-succostash/internal/core/service"
	"laughing-succostash/internal/handlers"
	"laughing-succostash/internal/repositories"
	"laughing-succostash/internal/validator"
	"log"
)

func init() {
	err := env.Load()
	if err != nil {
		log.Fatalf("Erro loading .env file")
	}
}

func main() {

	db := database.NewDatabase().Connect()

	repoUser := repositories.NewUserRepository(db)
	serviceUser := service.NewUserService(repoUser)
	validatorUser := validator.NewUserValidator()
	handlerUser := handlers.NewUserHttpHandler(serviceUser, validatorUser)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userHandler := e.Group("/user")
	userHandler.POST("", handlerUser.Create)
	userHandler.GET("", handlerUser.FindAll)
	userHandler.GET("/:id", handlerUser.Get)
	userHandler.DELETE("/:id", handlerUser.Delete)

	e.Logger.Fatal(e.Start(":8080"))

}
