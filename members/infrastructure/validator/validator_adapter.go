package validator

import (
	"errors"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslation "github.com/go-playground/validator/v10/translations/en"
)

type ValidatorAdapter struct {
	validator  *validator.Validate
	translator *ut.Translator
}

func NewValidatorAdapter() *ValidatorAdapter {
	validator := validator.New()

	en := en.New()
	translator, _ := ut.New(en, en).GetTranslator("en")
	enTranslation.RegisterDefaultTranslations(validator, translator)

	return &ValidatorAdapter{
		validator:  validator,
		translator: &translator,
	}
}

func (v *ValidatorAdapter) Validate(content interface{}) error {
	if err := v.validator.Struct(content); err != nil {
		return v.simplifyErrorsMessages(err)
	}
	return nil
}

func (v *ValidatorAdapter) simplifyErrorsMessages(err error) error {
	var errorsList []string
	for _, err := range err.(validator.ValidationErrors) {
		errorsList = append(errorsList, err.Translate(*v.translator))
	}
	return errors.New(strings.Join(errorsList, "; "))
}
