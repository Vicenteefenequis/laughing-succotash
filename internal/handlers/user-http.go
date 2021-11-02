package handlers

import (
	"github.com/labstack/echo/v4"
	"laughing-succostash/internal/core/domain"
	"laughing-succostash/internal/core/ports/service"
	validator_port "laughing-succostash/internal/core/ports/validator"
	"net/http"
)

type UserHTTPHandler struct {
	userService service_port.User
	validator   validator_port.UserValidator
}

func NewUserHttpHandler(bankService service_port.User, validator validator_port.UserValidator) *UserHTTPHandler {
	return &UserHTTPHandler{
		userService: bankService,
		validator:   validator,
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

	errorValidators := h.validator.Validate(u)

	if len(errorValidators) != 0 {
		c.JSON(http.StatusBadRequest, buildMessage("errors", errorValidators))
		return nil
	}

	user, err := h.userService.Create(*u)

	if err != nil {
		return err
	}

	c.JSON(http.StatusCreated, buildMessage("data", user))

	return nil
}

func (h *UserHTTPHandler) Delete(c echo.Context) error {
	err := h.userService.Delete(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, buildMessage("error", err.Error()))
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *UserHTTPHandler) FindAll(c echo.Context) error {
	users, err := h.userService.FindAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, buildMessage("err", err.Error()))
		return err
	}

	return c.JSON(http.StatusOK, buildMessage("data", users))
}
