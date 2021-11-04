package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"laughing-succostash/internal/core/service"
	"laughing-succostash/internal/handlers"
	"laughing-succostash/internal/repositories"
)

type User struct {
	e  *echo.Echo
	db *gorm.DB
}

func NewUserRouter(db *gorm.DB, e *echo.Echo) *User {
	return &User{
		db: db,
		e:  e,
	}
}

func (u *User) Router() {

	repoUser := repositories.NewUserRepository(u.db)
	serviceUser := service.NewUserService(repoUser)
	handlerUser := handlers.NewUserHttpHandler(serviceUser)

	userRoutes := u.e.Group("/user")
	userRoutes.POST("", handlerUser.Create)
	userRoutes.GET("", handlerUser.FindAll)
	userRoutes.GET("/:id", handlerUser.Get)
	userRoutes.DELETE("/:id", handlerUser.Delete)
	userRoutes.PUT("", handlerUser.Update)
}
