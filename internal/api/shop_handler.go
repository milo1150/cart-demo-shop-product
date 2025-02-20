package api

import (
	"minicart/internal/schemas"
	"minicart/internal/services"
	"minicart/internal/types"
	"minicart/internal/utils"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func CreateShopHandler(c echo.Context, appState *types.AppState) error {
	payload := &schemas.CreateShop{}

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	validate := validator.New()
	if err := validate.Struct(payload); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := utils.TranslateErrors(validationErrors)
		return c.JSON(http.StatusBadRequest, errorMessages)
	}

	shopService := services.ShopService{DB: appState.DB}
	if err := shopService.CreateShop(payload); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, http.StatusOK)
}
