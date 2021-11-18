package validator

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

type Validator struct {
	validator *validator.Validate
	trans     *ut.Translator
}

func NewValidator() *Validator {
	en := en.New()
	uni = ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	validator := &Validator{
		validator: validator.New(),
		trans:     &trans,
	}
	validator.RegisterTranslation()

	return validator
}

func (v *Validator) RegisterTranslation() error {
	err := en_translations.RegisterDefaultTranslations(v.validator, *v.trans)
	if err != nil {
		return err
	}
	err = v.RegisterTagRequired()
	if err != nil {
		return err
	}
	return nil
}

func (v *Validator) RegisterTagRequired() error {
	err := v.validator.RegisterTranslation("required", *v.trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} must have a value!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})
	if err != nil {
		return err
	}

	return nil
}
