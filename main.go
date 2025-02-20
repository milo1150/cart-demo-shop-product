package main

import (
	"minicart/internal/database"
	"minicart/internal/loader"
	"minicart/internal/routes"
	"minicart/internal/types"

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

	// Creates an instance of Echo.
	e := echo.New()

	// Init Route
	routes.RegisterAppRoutes(e, appState)

	// Start Server
	e.Logger.Fatal(e.Start(":1323"))
}
