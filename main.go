package main

import (
	"context"
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

	// Create a root context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Init zap logger
	logger := middlewares.InitializeZapLogger()

	// Connect postgres database
	gormDB := database.ConnectPostgresDatabase()

	// Migrate postgres
	database.RunAutoMigrate(gormDB)
	database.RunMigrate(gormDB)

	// Connect Minio
	minio := database.ConnectMinioDatabase()
	database.CreateBucket(minio, ctx, "product-image")

	// Init product table
	productLoader := loader.ProductLoader{Client: minio, Ctx: ctx, Log: logger}
	productLoader.InitializeProductDatas()

	// Global state
	appState := &types.AppState{
		DB:    gormDB,
		Minio: minio,
		Log:   logger,
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
