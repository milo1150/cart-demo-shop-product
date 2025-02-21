package routes

import (
	"minicart/internal/api"
	"minicart/internal/types"

	"github.com/labstack/echo/v4"
)

func ProductRoutes(e *echo.Echo, appState *types.AppState) {
	productGroup := e.Group("/product")

	productGroup.POST("/create", func(c echo.Context) error {
		return api.CreateProductHandler(c, appState)
	})
}
