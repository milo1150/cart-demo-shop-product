package routes

import (
	"shop-product-service/internal/api"
	"shop-product-service/internal/types"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Echo, appState *types.AppState) {
	productGroup := e.Group("/product")

	productGroup.POST("/create", func(c echo.Context) error {
		return api.CreateProductHandler(c, appState)
	})

	productGroup.POST("/generate-random-product", func(c echo.Context) error {
		return api.GenerateRandomProductHandler(c, appState)
	})
}
