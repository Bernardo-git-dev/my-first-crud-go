package validation

import (
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"

	"github.com/Bernardo-git-dev/my-first-crud-go/src/configuration/rest_err"
)

var (
	Validate          = validator.New()
	transl            ut.Translator
	jsonErr           *json.UnmarshalTypeError
	jsonValidationErr validator.ValidationErrors
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		un := ut.New(en, en)

		transl, _ = un.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validationErr error) *rest_err.RestErr {
	if errors.As(validationErr, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type")
	} else if errors.As(validationErr, &jsonValidationErr) {
		var causes []rest_err.Causes

		for _, e := range validationErr.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}
			causes = append(causes, cause)
		}

		return rest_err.NewBadRequestValidationError("Some fields are invalid", causes)
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields!")
	}
}
