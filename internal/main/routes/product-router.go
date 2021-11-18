package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"laughing-succostash/internal/core/service"
	"laughing-succostash/internal/handlers"
	"laughing-succostash/internal/repositories"
	"laughing-succostash/internal/validator"
)

type Product struct {
	e  *echo.Echo
	db *gorm.DB
}

func NewProductRouter(routes *Routes) *Product {
	return &Product{
		db: routes.db,
		e:  routes.e,
	}
}

func (u *Product) Router() {
	repoProduct := repositories.NewProductRepository(u.db)
	serviceProduct := service.NewProductService(repoProduct)
	validatorProduct := validator.NewProductValidator()
	handlerProduct := handlers.NewProductHttpHandler(serviceProduct, validatorProduct)

	userRoutes := u.e.Group("/product")
	userRoutes.POST("", handlerProduct.Create)
	userRoutes.GET("", handlerProduct.Find)
	userRoutes.DELETE("/:id", handlerProduct.Delete)
	userRoutes.PUT("", handlerProduct.Update)

}
