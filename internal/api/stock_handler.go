package api

import (
	"net/http"
	"shop-product-service/internal/schemas"
	"shop-product-service/internal/services"
	"shop-product-service/internal/types"
	"shop-product-service/internal/validators"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	cartpkg "github.com/milo1150/cart-demo-pkg/pkg"
)

func UpdateProductStockHandler(c echo.Context, appState *types.AppState) error {
	payload := &schemas.UpdateProductStockSlicesPayload{}
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	// Validate payload
	validate := validator.New()
	validate.RegisterValidation("stock_action", validators.StockActionValidator)
	if err := cartpkg.ValidateJsonPayload(validate, payload); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	stockService := services.StockService{DB: appState.DB}

	result, err := stockService.UpdateProductStock(payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	return c.JSON(http.StatusOK, result)
}
