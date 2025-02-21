package routes

import (
	"minicart/internal/types"

	"github.com/labstack/echo/v4"
)

func RegisterAppRoutes(e *echo.Echo, appState *types.AppState) {
	ShopRoutes(e, appState)
	ProductRoutes(e, appState)
	StockRoutes(e, appState)
}
