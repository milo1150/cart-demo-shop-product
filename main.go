package main

import (
	"context"
	"os"
	"shop-product-service/internal/database"
	"shop-product-service/internal/grpc"
	"shop-product-service/internal/loaders"
	"shop-product-service/internal/middlewares"
	"shop-product-service/internal/routes"
	"shop-product-service/internal/types"

	"github.com/labstack/echo/v4"
)

func main() {
	// Load ENV
	loaders.LoadENV()

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

	// Create Private bucket
	minioApiURL := os.Getenv("MINIO_API_URL") // FIXME: move and validate struct
	minioClient := database.MinIO{Client: minio, Context: ctx, ApiURL: minioApiURL, Log: logger}

	// Init Minio product images
	productMinioLoader := loaders.ProductMinIOLoader{Log: logger, Client: minio, Ctx: ctx}
	productMinioLoader.InitializeProductData("public-bucket", &minioClient)

	// Init Shop table
	shopPgLoader := loaders.ShopPgLoader{Ctx: ctx, Log: logger, DB: gormDB}
	shopPgLoader.InitializeShopData()

	// Init Product table
	productPgLoader := loaders.ProductPgLoader{Log: logger, DB: gormDB}
	productPgLoader.InitializeProductData(&minioClient, "public-bucket")

	// Global state
	appState := &types.AppState{
		DB:    gormDB,
		Log:   logger,
		Minio: minio,
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
