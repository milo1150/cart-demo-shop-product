package main

import (
	"minicart/src/database"
	"minicart/src/loader"

	"github.com/labstack/echo/v4"
)

func main() {
	// Load ENV
	loader.LoadEnv()

	// Database handler
	database.ConnectDatabase()

	// Start Server
	e := echo.New()
	e.Logger.Fatal(e.Start(":1323"))
}
