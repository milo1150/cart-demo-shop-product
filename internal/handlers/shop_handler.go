package handlers

import (
	"minicart/internal/schemas"
	"minicart/internal/types"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateShopHandler(c echo.Context, appState *types.AppState) error {
	payload := schemas.CreateShop{}
	c.Bind(&payload)
	// result := repository.CreateShop(appState.DB)
	return c.JSON(http.StatusOK, http.StatusOK)
}
