package main

import (
	"minicart/src/database"
	"minicart/src/loader"
	"minicart/src/route"
	"minicart/src/types"

	"github.com/labstack/echo/v4"
)

func main() {
	// Load ENV
	loader.LoadEnv()

	// Database handler
	db := database.ConnectDatabase()
	database.RunAutoMigrate(db)

	// Global state
	appState := &types.AppState{
		DB: db,
	}

	e := echo.New()

	// Main route
	route := &route.AppRoute{Echo: e, AppState: appState}
	route.RegisterAppRoutes()

	// Start Server
	e.Logger.Fatal(e.Start(":1323"))
}
