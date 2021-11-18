package handlers

import (
	"github.com/labstack/echo/v4"
	"laughing-succostash/internal/core/domain"
	service_port "laughing-succostash/internal/core/ports/service"
	validator_port "laughing-succostash/internal/core/ports/validator"
	"net/http"
)

type CategoryHTTPHandler struct {
	categoryService service_port.Category
	validator       validator_port.Validator
}

func NewCategoryHttpHandler(categoryService service_port.Category, validator validator_port.Validator) *CategoryHTTPHandler {
	return &CategoryHTTPHandler{
		categoryService: categoryService,
		validator:       validator,
	}
}

func (h *CategoryHTTPHandler) Create(c echo.Context) error {
	category := new(domain.Category)

	if err := c.Bind(category); err != nil {
		c.JSON(http.StatusBadRequest, buildMessage("error", err.Error()))
		return err
	}
	errorsValidator := h.validator.Validate(*category)

	if len(errorsValidator) != 0 {
		c.JSON(http.StatusBadRequest, buildMessage("errors", errorsValidator))
		return nil
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

func (h *CategoryHTTPHandler) Find(c echo.Context) error {
	names := getIdsParam(c.QueryParam("names"))

	limit, offset := getPaginationParam(c)

	categories, err := h.categoryService.Find(names, limit, offset)

	if err != nil {
		c.JSON(http.StatusBadRequest, buildMessage("err", err.Error()))
		return err
	}

	setPagination(c, limit > len(categories))

	return c.JSON(http.StatusOK, buildMessage("data", categories))
}

func (h *CategoryHTTPHandler) Update(c echo.Context) error {
	u := new(domain.Category)

	if err := c.Bind(u); err != nil {
		c.JSON(http.StatusBadRequest, buildMessage("error", err.Error()))
		return err
	}

	errorsValidator := h.validator.Validate(*u)

	if len(errorsValidator) != 0 {
		c.JSON(http.StatusBadRequest, buildMessage("errors", errorsValidator))
		return nil
	}

	_category, err := h.categoryService.Update(*u)

	if err != nil {
		return err
	}

	c.JSON(http.StatusCreated, buildMessage("data", _category))

	return err
}
