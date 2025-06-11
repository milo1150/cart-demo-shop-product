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

	// Init MinIO
	minio := database.ConnectMinioDatabase()
	minioApiURL := os.Getenv("MINIO_API_URL")
	publicBucketName := os.Getenv("MINIO_PUBLIC_BUCKET_NAME")
	minioClient := database.MinIO{Client: minio, Context: ctx, ApiURL: minioApiURL, Log: logger}

	// Init MinIO Bucket
	database.CreatePublicBucket(minioClient.Client, publicBucketName)

	// Init Minio product images
	productMinioLoader := loaders.ProductMinIOLoader{Log: logger, Client: minio, Ctx: ctx}
	productMinioLoader.InitializeProductData(publicBucketName, &minioClient)

	// Init Shop table
	shopPgLoader := loaders.ShopPgLoader{Ctx: ctx, Log: logger, DB: gormDB}
	shopPgLoader.InitializeShopData()

	// Init Product table
	productPgLoader := loaders.ProductPgLoader{Log: logger, DB: gormDB}
	productPgLoader.InitializeProductData(&minioClient, publicBucketName)

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
