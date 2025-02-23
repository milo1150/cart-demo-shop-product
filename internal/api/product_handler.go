package api

import (
	"net/http"
	"shop-product-service/internal/schemas"
	"shop-product-service/internal/services"
	"shop-product-service/internal/types"
	"shop-product-service/internal/utils"
	"shop-product-service/internal/validators"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func CreateProductHandler(c echo.Context, appState *types.AppState) error {
	payload := &schemas.CreateProductSchema{}
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, utils.GetSimpleErrorMessage(err.Error()))
	}

	validate := validator.New()
	if err := validators.ValidateJsonPayload(validate, payload); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	productService := services.ProductService{DB: appState.DB}
	if err := productService.CreateProduct(payload); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.GetSimpleErrorMessage(err.Error()))
	}

	return c.JSON(http.StatusOK, http.StatusOK)
}
