package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"laughing-succostash/internal/core/service"
	"laughing-succostash/internal/handlers"
	"laughing-succostash/internal/repositories"
	"laughing-succostash/internal/validator"
)

type User struct {
	e  *echo.Echo
	db *gorm.DB
}

func NewUserRouter(routes *Routes) *User {
	return &User{
		db: routes.db,
		e:  routes.e,
	}
}

func (u *User) Router() {
	repoUser := repositories.NewUserRepository(u.db)
	serviceUser := service.NewUserService(repoUser)
	validatorUser := validator.NewUserValidator()
	handlerUser := handlers.NewUserHttpHandler(serviceUser, validatorUser)

	userRoutes := u.e.Group("/user")
	userRoutes.POST("", handlerUser.Create)
	userRoutes.GET("", handlerUser.Find)
	userRoutes.DELETE("/:id", handlerUser.Delete)
	userRoutes.PUT("", handlerUser.Update)
}
