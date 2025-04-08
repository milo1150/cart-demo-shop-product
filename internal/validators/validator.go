package validators

import (
	"shop-product-service/internal/enums"

	"github.com/go-playground/validator/v10"
)

func StockActionValidator(fl validator.FieldLevel) bool {
	action, ok := fl.Field().Interface().(enums.StockAction)
	if !ok {
		return false
	}
	return action.IsValid()
}
