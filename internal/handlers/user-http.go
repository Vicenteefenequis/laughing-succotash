package handlers

import (
	"github.com/labstack/echo/v4"
	"laughing-succostash/internal/core/domain"
	"laughing-succostash/internal/core/ports/service"
	"net/http"
)

type UserHTTPHandler struct {
	userService service_port.User
}

func NewHttpHandler(bankService service_port.User) *UserHTTPHandler {
	return &UserHTTPHandler{
		userService: bankService,
	}
}

func (h *UserHTTPHandler) Get(c echo.Context) error {
	user, err := h.userService.FindOne(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, buildMessage("error", err.Error()))
		return err
	}

	c.JSON(http.StatusOK, user)
	return nil
}

func (h *UserHTTPHandler) Create(c echo.Context) error {
	u := new(domain.User)

	if err := c.Bind(u); err != nil {
		c.JSON(http.StatusBadRequest, buildMessage("error", err.Error()))
		return err
	}

	user, err := h.userService.Create(*u)

	if err != nil {
		return err
	}

	c.JSON(http.StatusCreated, buildMessage("data", user))

	return nil

}
