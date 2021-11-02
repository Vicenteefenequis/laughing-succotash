package validator

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	"laughing-succostash/internal/core/domain"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

type UserInput struct {
	Name string      `json:"name" validate:"required"`
	Type domain.Type `json:"type" validate:"required"`
}

type UserValidator struct {
	validator *validator.Validate
	trans     *ut.Translator
}

func NewUserValidator() *UserValidator {
	en := en.New()
	uni = ut.New(en, en)
	trans, _ := uni.GetTranslator("en")

	return &UserValidator{
		validator: validator.New(),
		trans:     &trans,
	}
}

func (u *UserValidator) Validate(user *domain.User) []string {
	err := u.RegisterTranslation()

	if err != nil {
		return []string{err.Error()}
	}

	if !isTypeValid(user.Type) {
		return []string{"Type must be a client or store"}
	}

	_user := &UserInput{
		Name: user.Name,
		Type: user.Type,
	}

	err = u.validator.Struct(_user)

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

func (u *UserValidator) RegisterTranslation() error {

	err := en_translations.RegisterDefaultTranslations(u.validator, *u.trans)
	if err != nil {
		return err
	}

	err = u.RegisterTagRequired()
	if err != nil {
		return err
	}

	return nil
}

func (u *UserValidator) RegisterTagRequired() error {
	err := u.validator.RegisterTranslation("required", *u.trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} must have a value!", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})
	if err != nil {
		return err
	}

	return nil
}

func isTypeValid(_type domain.Type) bool {
	if _type == "client" || _type == "store" {
		return true
	}
	return false
}
