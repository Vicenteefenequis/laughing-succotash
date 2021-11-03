package main

import (
	env "github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"laughing-succostash/internal/core/database"
	"laughing-succostash/internal/main/routes"
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
	e := echo.New()

	userRouter := routes.NewUserRouter(db, e)
	userRouter.UserRouter()

	categoryRouter := routes.NewCategoryRouter(db, e)
	categoryRouter.CategoryRouter()

	e.Logger.Fatal(e.Start(":8080"))

}
