package api

import (
	"net/http"
	"shop-product-service/internal/schemas"
	"shop-product-service/internal/services"
	"shop-product-service/internal/types"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	cartpkg "github.com/milo1150/cart-demo-pkg/pkg"
)

func UpdateProductStockHandler(c echo.Context, appState *types.AppState) error {
	payload := &schemas.UpdateProductStockSlicesPayload{}
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	validate := validator.New()
	if err := cartpkg.ValidateJsonPayload(validate, payload); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	stockService := services.StockService{DB: appState.DB}

	result, err := stockService.UpdateProductStockService(payload)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	return c.JSON(http.StatusOK, result)
}
