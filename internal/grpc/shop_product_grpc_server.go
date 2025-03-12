package grpc

import (
	"context"
	"log"
	"net"
	"shop-product-service/internal/repositories"
	"shop-product-service/internal/types"

	pb "github.com/milo1150/cart-demo-proto/pkg/shop_product"
	"google.golang.org/grpc"
)

type ShopProductServer struct {
	pb.UnimplementedShopProductServiceServer

	AppState *types.AppState
}

func StartShopProductGRPCServer(appState *types.AppState) {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterShopProductServiceServer(grpcServer, &ShopProductServer{AppState: appState})

	log.Println("gRPC server is running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve ShopProduct gRPC Server: %v", err)
	}
}

func (s *ShopProductServer) GetProduct(context.Context, *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	return &pb.GetProductResponse{Id: 1, ProductName: "test grpc", Price: 50, Stock: 100, ShopId: 99}, nil
}

func (s *ShopProductServer) ProductExists(_ context.Context, payload *pb.CheckProductRequest) (*pb.CheckProductReponse, error) {
	productRepository := repositories.ProductRepository{DB: s.AppState.DB}

	result, err := productRepository.VerifyIsProductExistsByID(uint(payload.ProductId))
	if err != nil {
		return nil, err
	}

	return &pb.CheckProductReponse{IsExists: result}, nil
}
