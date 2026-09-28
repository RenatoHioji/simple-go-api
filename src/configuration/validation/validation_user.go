package validation

import (
	"encoding/json"
	"errors"

	"github.com/RenatoHioji/simple-go-api/src/configuration/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	translat ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		un := ut.New(en, en)
		transl, _ := un.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validation_err error) *rest_err.RestErr {
	var json_err *json.UnmarshalTypeError
	var json_validation_err validator.ValidationErrors

	if errors.As(validation_err, &json_err) {
		return rest_err.NewBadRequestErr("Invalid field type")
	}

	if errors.As(validation_err, &json_validation_err) {
		errorsCauses := []rest_err.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Message: e.Translate(translat),
				Field:   e.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}

		return rest_err.NewBadRequestValidationErr("Some fields are invalid", errorsCauses)
	}

	return rest_err.NewBadRequestErr("Error trying to convert fields")
}
