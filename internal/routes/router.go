package routes

import (
	"minicart/internal/handlers"
	"minicart/internal/types"

	"github.com/labstack/echo/v4"
)

type AppRoute struct {
	Echo     *echo.Echo
	AppState *types.AppState
}

func (r *AppRoute) RegisterAppRoutes() {
	r.shopRoutes()
}

func (r *AppRoute) shopRoutes() {
	shopGroup := r.Echo.Group("/shop")
	shopGroup.POST("/create", func(c echo.Context) error {
		return handlers.CreateShopHandler(c, r.AppState)
	})
}
