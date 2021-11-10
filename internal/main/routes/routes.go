package routes

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Routes struct {
	e  *echo.Echo
	db *gorm.DB
}

func New(e *echo.Echo, db *gorm.DB) *Routes {
	return &Routes{
		e:  e,
		db: db,
	}
}

func (r *Routes) Register() {

	user := NewUserRouter(r)
	user.Router()

	category := NewCategoryRouter(r)
	category.Router()

	product := NewProductRouter(r)
	product.Router()
}
