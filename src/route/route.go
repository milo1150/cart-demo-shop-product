package route

import (
	"minicart/src/handler"
	"minicart/src/types"

	"github.com/labstack/echo/v4"
)

type AppRoute struct {
	Echo     *echo.Echo
	AppState *types.AppState
}

func (r *AppRoute) RegisterAppRoutes() {
	r.shopRoutes()
	r.couponRoutes()
}

func (r *AppRoute) shopRoutes() {
	shopGroup := r.Echo.Group("/shop")
	shopGroup.POST("/create", func(c echo.Context) error {
		return handler.CreateShopHandler(c, r.AppState)
	})
}

func (r *AppRoute) couponRoutes() {
	couponGroup := r.Echo.Group("/coupon")
	couponGroup.POST("/create", func(c echo.Context) error {
		return nil
	})
}
