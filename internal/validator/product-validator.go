package validator

import (
	"github.com/go-playground/validator/v10"
	"laughing-succostash/internal/core/domain"
)

type ProductInput struct {
	Name       string `json:"name" validate:"required"`
	Quantity   int    `json:"quantity" validate:"required"`
	Price      uint64 `json:"price" validate:"required"`
	CategoryId string `json:"category_id" validate:"required"`
}

type ProductValidator struct {
	Validator
}

func NewProductValidator() *ProductValidator {
	return &ProductValidator{
		*NewValidator(),
	}
}

func (p *ProductValidator) Validate(field interface{}) []string {
	product := field.(domain.Product)

	_product := &ProductInput{
		Name:       product.Name,
		Quantity:   product.Quantity,
		Price:      product.Price,
		CategoryId: product.CategoryId,
	}

	err := p.validator.Struct(_product)

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
