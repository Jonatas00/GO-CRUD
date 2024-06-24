package validation

import (
	"encoding/json"
	"errors"

	"github.com/Jonatas00/GO-CRUD/internal/config/restErr"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		trans, _ := unt.GetTranslator("en")
		en_translations.RegisterDefaultTranslations(val, trans)
	}
}

func ValidateUserError(validationError error) *restErr.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError *validator.ValidationErrors

	if errors.As(validationError, &jsonErr) {
		return restErr.NewBadRequestError("invalid field type")
	} else if errors.As(validationError, &jsonValidationError) {
		errorsCauses := []restErr.Causes{}

		for _, e := range validationError.(validator.ValidationErrors) {
			cause := restErr.Causes{
				Message: e.Translate(transl),
				Field:   e.Field(),
			}

			errorsCauses = append(errorsCauses, cause)
		}
		return restErr.NewBadRequestValidationError("Some fields are invalids", errorsCauses)
	} else {
		return restErr.NewBadRequestError("Error trying to convert fields")
	}
}
