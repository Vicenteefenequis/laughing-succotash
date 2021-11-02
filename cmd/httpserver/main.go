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

	repoUser := repositories.NewUserRepository(db)
	serviceUser := service.NewUserService(repoUser)
	handlerUser := handlers.NewUserHttpHandler(serviceUser)

	repoCategory := repositories.NewCategoryRepository(db)
	serviceCategory := service.NewCategoryService(repoCategory)
	handlerCategory := handlers.NewCategoryHttpHandler(serviceCategory)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userHandler := e.Group("/user")
	userHandler.POST("", handlerUser.Create)
	userHandler.GET("", handlerUser.FindAll)
	userHandler.GET("/:id", handlerUser.Get)
	userHandler.DELETE("/:id", handlerUser.Delete)
	userHandler.PUT("", handlerUser.Update)

	categoryHandler := e.Group("category")
	categoryHandler.POST("", handlerCategory.Create)
	categoryHandler.GET("", handlerCategory.FindAll)
	categoryHandler.GET("/:id", handlerCategory.Get)
	categoryHandler.DELETE("/:id", handlerCategory.Delete)
	categoryHandler.PUT("", handlerCategory.Update)

	e.Logger.Fatal(e.Start(":8080"))

}
