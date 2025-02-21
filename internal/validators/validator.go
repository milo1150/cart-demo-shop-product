package validators

import (
	"minicart/internal/utils"

	"github.com/go-playground/validator/v10"
)

func ValidateJsonPayload(validate *validator.Validate, payload interface{}) map[string]string {
	if err := validate.Struct(payload); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		return utils.TranslateErrors(validationErrors)
	}
	return nil
}
