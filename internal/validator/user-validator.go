package validator

import (
	"github.com/go-playground/validator/v10"
	"laughing-succostash/internal/core/domain"
)

type UserInput struct {
	Name string      `json:"name" validate:"required"`
	Type domain.Type `json:"type" validate:"required"`
}

type UserValidator struct {
	Validator
}

func NewUserValidator() *UserValidator {
	return &UserValidator{
		*NewValidator(),
	}
}

func (u *UserValidator) Validate(field interface{}) []string {
	user := field.(domain.User)

	if !u.IsTypeValid(user.Type) {
		return []string{"Type must be a client or store"}
	}

	_user := &UserInput{
		Name: user.Name,
		Type: user.Type,
	}

	err := u.validator.Struct(_user)

	var errorValidator []string

	if err != nil {

		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			errorValidator = append(errorValidator, e.Translate(*u.trans))
		}
		return errorValidator
	}

	return []string{}
}

func (u *UserValidator) IsTypeValid(_type domain.Type) bool {
	if _type == "client" || _type == "store" {
		return true
	}
	return false
}
