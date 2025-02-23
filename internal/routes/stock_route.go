package routes

import (
	"shop-product-service/internal/api"
	"shop-product-service/internal/types"

	"github.com/labstack/echo/v4"
)

func StockRoutes(e *echo.Echo, appState *types.AppState) {
	stockGroup := e.Group("/stock")

	stockGroup.POST("/update-product-stock", func(c echo.Context) error {
		return api.UpdateProductStockHandler(c, appState)
	})
}
