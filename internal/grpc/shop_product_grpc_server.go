package grpc

import (
	"context"
	"log"
	"net"

	pb "github.com/milo1150/cart-demo-proto/pkg/shop_product"
	"google.golang.org/grpc"
)

type ShopProductServer struct {
	pb.UnimplementedShopProductServiceServer
}

func (s *ShopProductServer) GetProduct(context.Context, *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	return &pb.GetProductResponse{Id: 1, ProductName: "test grpc", Price: 50, Stock: 100, ShopId: 99}, nil
}

func StartShopProductGRPCServer() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterShopProductServiceServer(grpcServer, &ShopProductServer{})

	log.Println("gRPC server is running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve ShopProduct gRPC Server: %v", err)
	}
}
