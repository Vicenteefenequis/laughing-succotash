package main

import (
	env "github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"laughing-succostash/internal/core/database"
	"laughing-succostash/internal/core/service"
	"laughing-succostash/internal/handlers"
	"laughing-succostash/internal/repositories"
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
	repo := repositories.NewUserRepository(db)
	service := service.NewUserService(repo)
	handler := handlers.NewHttpHandler(service)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userHandler := e.Group("/user")
	userHandler.POST("", handler.Create)
	userHandler.GET("/:id", handler.Get)

	e.Logger.Fatal(e.Start(":8080"))

}
