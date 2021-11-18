package validator

import (
	"github.com/go-playground/validator/v10"
	"laughing-succostash/internal/core/domain"
)

type CategoryInput struct {
	Name string `json:"name" validate:"required"`
}

type CategoryValidator struct {
	Validator
}

func NewCategoryValidator() *CategoryValidator {
	return &CategoryValidator{
		*NewValidator(),
	}
}

func (p *CategoryValidator) Validate(field interface{}) []string {
	category := field.(domain.Category)

	_category := &CategoryInput{
		Name: category.Name,
	}

	err := p.validator.Struct(_category)

	var errorValidator []string

	if err != nil {

		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			errorValidator = append(errorValidator, e.Translate(*p.trans))
		}
		return errorValidator
	}

	return []string{}
}
