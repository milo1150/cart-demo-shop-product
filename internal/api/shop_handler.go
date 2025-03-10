package api

import (
	"net/http"
	"shop-product-service/internal/schemas"
	"shop-product-service/internal/services"
	"shop-product-service/internal/types"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	cartpkg "github.com/milo1150/cart-demo-pkg/pkg"
)

func CreateShopHandler(c echo.Context, appState *types.AppState) error {
	payload := &schemas.CreateShop{}

	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	validate := validator.New()
	if err := cartpkg.ValidateJsonPayload(validate, payload); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	shopService := services.ShopService{DB: appState.DB}
	if err := shopService.CreateShop(payload); err != nil {
		return c.JSON(http.StatusInternalServerError, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	return c.JSON(http.StatusOK, http.StatusOK)
}

func GetShopDetailHandler(c echo.Context, appState *types.AppState) error {
	shopUuid, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, cartpkg.GetSimpleErrorMessage("Invalid shop id"))
	}

	shopService := services.ShopService{DB: appState.DB}

	shop, err := shopService.GetShopDetail(shopUuid)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, cartpkg.GetSimpleErrorMessage(err.Error()))
	}

	return c.JSON(http.StatusOK, shop)
}
