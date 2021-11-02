package handlers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"laughing-succostash/internal/core/domain"
	service_port "laughing-succostash/internal/core/ports/service"
	"net/http"
)

type CategoryHTTPHandler struct {
	categoryService service_port.Category
}

func NewCategoryHttpHandler(categoryService service_port.Category) *CategoryHTTPHandler {
	return &CategoryHTTPHandler{
		categoryService: categoryService,
	}
}

func (h *CategoryHTTPHandler) Get(c echo.Context) error {
	user, err := h.categoryService.FindOne(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, buildMessage("error", err.Error()))
		return err
	}

	c.JSON(http.StatusOK, user)
	return nil
}

func (h *CategoryHTTPHandler) Create(c echo.Context) error {
	category := new(domain.Category)

	if err := c.Bind(category); err != nil {
		c.JSON(http.StatusBadRequest, buildMessage("error", err.Error()))
		return err
	}

	_category, err := h.categoryService.Create(*category)

	if err != nil {
		return err
	}

	c.JSON(http.StatusCreated, buildMessage("data", _category))

	return err
}

func (h *CategoryHTTPHandler) Delete(c echo.Context) error {
	err := h.categoryService.Delete(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, buildMessage("error", err.Error()))
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *CategoryHTTPHandler) FindAll(c echo.Context) error {
	categories, err := h.categoryService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, buildMessage("err", err.Error()))
		return err
	}

	return c.JSON(http.StatusOK, buildMessage("data", categories))
}

func (h *CategoryHTTPHandler) Update(c echo.Context) error {
	u := new(domain.Category)

	if err := c.Bind(u); err != nil {
		c.JSON(http.StatusBadRequest, buildMessage("error", err.Error()))
		return err
	}

	fmt.Println(*u)

	_category, err := h.categoryService.Update(*u)

	if err != nil {
		return err
	}

	c.JSON(http.StatusCreated, buildMessage("data", _category))

	return err
}
