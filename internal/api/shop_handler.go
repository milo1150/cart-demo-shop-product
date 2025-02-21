package api

import (
	"minicart/internal/schemas"
	"minicart/internal/services"
	"minicart/internal/types"
	"minicart/internal/utils"
	"minicart/internal/validators"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func CreateShopHandler(c echo.Context, appState *types.AppState) error {
	payload := &schemas.CreateShop{}

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, utils.GetSimpleErrorMessage(err.Error()))
	}

	validate := validator.New()
	if err := validators.ValidateJsonPayload(validate, payload); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	shopService := services.ShopService{DB: appState.DB}
	if err := shopService.CreateShop(payload); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.GetSimpleErrorMessage(err.Error()))
	}

	return c.JSON(http.StatusOK, http.StatusOK)
}

func GetShopDetailHandler(c echo.Context, appState *types.AppState) error {
	shopUuid, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.GetSimpleErrorMessage("Invalid shop id"))
	}

	shopService := services.ShopService{DB: appState.DB}

	shop, err := shopService.GetShopDetail(shopUuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.GetSimpleErrorMessage(err.Error()))
	}

	return c.JSON(http.StatusOK, shop)
}
