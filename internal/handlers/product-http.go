package handlers

import (
	"github.com/labstack/echo/v4"
	"laughing-succostash/internal/core/domain"
	service_port "laughing-succostash/internal/core/ports/service"
	"net/http"
)

type ProductHttpHandler struct {
	productService service_port.Product
}

func NewProductHttpHandler(productService service_port.Product) *ProductHttpHandler {
	return &ProductHttpHandler{
		productService: productService,
	}
}

func (h *ProductHttpHandler) Create(c echo.Context) error {
	u := new(domain.Product)

	if err := c.Bind(u); err != nil {
		c.JSON(http.StatusBadRequest, buildMessage("error", err.Error()))
		return err
	}

	product, err := h.productService.Create(*u)

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

	_product, err := h.productService.Update(*u)

	if err != nil {
		return err
	}

	c.JSON(http.StatusCreated, buildMessage("data", _product))

	return err
}
