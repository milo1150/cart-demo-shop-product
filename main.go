package main

import (
	"shop-product-service/internal/database"
	"shop-product-service/internal/grpc"
	"shop-product-service/internal/loader"
	"shop-product-service/internal/middlewares"
	"shop-product-service/internal/routes"
	"shop-product-service/internal/types"

	"github.com/labstack/echo/v4"
)

func main() {
	// Load ENV
	loader.LoadEnv()

	// Database handler
	db := database.ConnectDatabase()
	database.RunMigrate(db)

	// Global state
	appState := &types.AppState{
		DB: db,
	}

	// Creates an instance of Echo.
	e := echo.New()

	// Middlewares
	middlewares.RegisterMiddlewares(e)

	// Init Route
	routes.RegisterAppRoutes(e, appState)

	// gRPC Server
	go grpc.StartShopProductGRPCServer(appState)

	// Start Server
	go e.Logger.Fatal(e.Start(":1323"))
}
