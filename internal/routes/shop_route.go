package routes

import (
	"shop-product-service/internal/api"
	"shop-product-service/internal/types"

	"github.com/labstack/echo/v4"
)

func ShopRoutes(e *echo.Echo, appState *types.AppState) {
	shopGroup := e.Group("/shop")

	shopGroup.POST("/create", func(c echo.Context) error {
		return api.CreateShopHandler(c, appState)
	})

	shopGroup.GET("/:uuid", func(c echo.Context) error {
		return api.GetShopDetailHandler(c, appState)
	})
}
