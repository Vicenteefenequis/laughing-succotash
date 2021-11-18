package handlers

import (
	"github.com/labstack/echo/v4"
	"laughing-succostash/internal/core/domain"
	service_port "laughing-succostash/internal/core/ports/service"
	validator_port "laughing-succostash/internal/core/ports/validator"
	"net/http"
)

type ProductHttpHandler struct {
	productService service_port.Product
	validator      validator_port.Validator
}

func NewProductHttpHandler(productService service_port.Product, validator validator_port.Validator) *ProductHttpHandler {
	return &ProductHttpHandler{
		productService: productService,
		validator:      validator,
	}
}

func (h *ProductHttpHandler) Create(c echo.Context) error {
	p := new(domain.Product)

	if err := c.Bind(p); err != nil {
		c.JSON(http.StatusBadRequest, buildMessage("error", err.Error()))
		return err
	}

	errorsValidator := h.validator.Validate(*p)

	if len(errorsValidator) != 0 {
		c.JSON(http.StatusBadRequest, buildMessage("errors", errorsValidator))
		return nil
	}

	product, err := h.productService.Create(*p)

	if err != nil {
		return err
	}

	c.JSON(http.StatusCreated, buildMessage("data", product))

	return err
}

func (h *ProductHttpHandler) Delete(c echo.Context) error {
	err := h.productService.Delete(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, buildMessage("error", err.Error()))
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *ProductHttpHandler) Find(c echo.Context) error {
	ids := getIdsParam(c.QueryParam("ids"))

	limit, offset := getPaginationParam(c)

	products, err := h.productService.Find(ids, limit, offset)

	if err != nil {
		c.JSON(http.StatusBadRequest, buildMessage("err", err.Error()))
		return err
	}

	setPagination(c, limit > len(products))

	return c.JSON(http.StatusOK, buildMessage("data", products))
}

func (h *ProductHttpHandler) Update(c echo.Context) error {
	u := new(domain.Product)

	if err := c.Bind(u); err != nil {
		c.JSON(http.StatusBadRequest, buildMessage("error", err.Error()))
		return err
	}

	errorsValidator := h.validator.Validate(*u)

	if len(errorsValidator) != 0 {
		c.JSON(http.StatusBadRequest, buildMessage("errors", errorsValidator))
		return nil
	}

	_product, err := h.productService.Update(*u)

	if err != nil {
		return err
	}

	c.JSON(http.StatusCreated, buildMessage("data", _product))

	return err
}
