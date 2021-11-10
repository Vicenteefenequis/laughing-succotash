package handlers

import (
	"github.com/labstack/echo/v4"
	"laughing-succostash/internal/core/domain"
	"laughing-succostash/internal/core/ports/service"
	validator_port "laughing-succostash/internal/core/ports/validator"
	"laughing-succostash/internal/validator"
	"net/http"
)

type UserHTTPHandler struct {
	userService service_port.User
	validator   validator_port.Validator
}

func NewUserHttpHandler(bankService service_port.User, validator validator_port.Validator) *UserHTTPHandler {
	return &UserHTTPHandler{
		userService: bankService,
		validator:   validator,
	}
}

func (h *UserHTTPHandler) Create(c echo.Context) error {
	u := new(domain.User)

	if err := c.Bind(u); err != nil {
		c.JSON(http.StatusBadRequest, buildMessage("error", err.Error()))
		return err
	}

	errorValidators := h.validator.Validate(*u)

	if len(errorValidators) != 0 {
		c.JSON(http.StatusBadRequest, buildMessage("errors", errorValidators))
		return nil
	}

	user, err := h.userService.Create(*u)

	if err != nil {
		return err
	}

	c.JSON(http.StatusCreated, buildMessage("data", user))

	return err
}

func (h *UserHTTPHandler) Delete(c echo.Context) error {
	err := h.userService.Delete(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, buildMessage("error", err.Error()))
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *UserHTTPHandler) Find(c echo.Context) error {
	ids := getIdsParam(c.QueryParam("ids"))

	limit, offset := getPaginationParam(c)

	users, err := h.userService.Find(ids, limit, offset)

	if err != nil {
		c.JSON(http.StatusBadRequest, buildMessage("err", err.Error()))
		return err
	}

	setPagination(c, limit > len(users))

	return c.JSON(http.StatusOK, buildMessage("data", users))
}

func (h *UserHTTPHandler) Update(c echo.Context) error {
	u := new(domain.User)

	if err := c.Bind(u); err != nil {
		c.JSON(http.StatusBadRequest, buildMessage("error", err.Error()))
		return err
	}

	validator := validator.NewUserValidator()

	if u.Type != "" && !validator.IsTypeValid(u.Type) {
		return c.JSON(http.StatusBadRequest, buildMessage("error", "Type must be a client or store"))
	}

	_user, err := h.userService.Update(*u)

	if err != nil {
		return err
	}

	c.JSON(http.StatusCreated, buildMessage("data", _user))

	return err
}
