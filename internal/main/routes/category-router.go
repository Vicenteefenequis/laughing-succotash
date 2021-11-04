package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"laughing-succostash/internal/core/service"
	"laughing-succostash/internal/handlers"
	"laughing-succostash/internal/repositories"
)

type Category struct {
	e  *echo.Echo
	db *gorm.DB
}

func NewCategoryRouter(db *gorm.DB, e *echo.Echo) *Category {
	return &Category{
		db: db,
		e:  e,
	}
}

func (u *Category) Router() {

	repoCategory := repositories.NewCategoryRepository(u.db)
	serviceCategory := service.NewCategoryService(repoCategory)
	handlerCategory := handlers.NewCategoryHttpHandler(serviceCategory)

	categoryRoutes := u.e.Group("category")
	categoryRoutes.POST("", handlerCategory.Create)
	categoryRoutes.GET("", handlerCategory.FindAll)
	categoryRoutes.GET("/:id", handlerCategory.Get)
	categoryRoutes.DELETE("/:id", handlerCategory.Delete)
	categoryRoutes.PUT("", handlerCategory.Update)
}
