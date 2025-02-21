package routes

import (
	"minicart/internal/api"
	"minicart/internal/types"

	"github.com/labstack/echo/v4"
)

func StockRoutes(e *echo.Echo, appState *types.AppState) {
	stockGroup := e.Group("/stock")

	stockGroup.POST("/update-product-stock", func(c echo.Context) error {
		return api.UpdateProductStockHandler(c, appState)
	})
}
