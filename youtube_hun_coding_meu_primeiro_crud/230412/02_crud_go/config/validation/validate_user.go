package validation

import (
	"crud_go/config/rest_err"
	"encoding/json"
	"errors"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	trans    ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		univTranslator := ut.New(en, en)
		trans, _ = univTranslator.GetTranslator("en")

		en_translation.RegisterDefaultTranslations(val, trans)
	}
}

func ValidateUserError(validationErr error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validationErr, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type")
	} else if errors.As(validationErr, &jsonValidationError) {
		errorCauses := []rest_err.Cause{}
		for _, e := range validationErr.(validator.ValidationErrors) {
			cause := rest_err.Cause{
				Message: e.Translate(trans),
				Field:   e.Field(),
			}

			errorCauses = append(errorCauses, cause)
		}
		return rest_err.NewBadRequestValidationError("Some fields are invalid", errorCauses)
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields.")
	}
}
