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

func (s *ShopProductServer) GetProduct(_ context.Context, payload *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	productRepository := repositories.ProductRepository{DB: s.AppState.DB}

	product, err := productRepository.FindProductByID(uint(payload.ProductId))
	if err != nil {
		return nil, err
	}

	shopDetail := &pb.ShopDetail{
		Id:   uint64(product.Shop.ID),
		Name: product.Shop.Name,
	}

	res := &pb.GetProductResponse{
		Id:          uint64(product.ID),
		Name:        product.Name,
		Description: product.Description,
		Price:       float64(product.Price),
		Stock:       uint64(product.Stock),
		Shop:        shopDetail,
		Image:       product.ImageUrl,
	}

	return res, nil
}

func (s *ShopProductServer) GetProducts(_ context.Context, payload *pb.GetProductsRequest) (*pb.GetProductsResponse, error) {
	productRepository := repositories.ProductRepository{DB: s.AppState.DB}
	log.Println("gRPC - GetProducts payload:", payload)

	products, err := productRepository.FindProductsByIDs(payload.ProductIds)
	if err != nil {
		return nil, err
	}
	log.Println("gRPC - GetProducts length:", len(*products))

	res := &pb.GetProductsResponse{Products: []*pb.GetProductResponse{}}
	for _, product := range *products {
		shopDetail := &pb.ShopDetail{
			Id:   uint64(product.Shop.ID),
			Name: product.Shop.Name,
		}

		productRes := &pb.GetProductResponse{
			Id:          uint64(product.ID),
			Name:        product.Name,
			Description: product.Description,
			Price:       float64(product.Price),
			Stock:       uint64(product.Stock),
			Shop:        shopDetail,
			Image:       product.ImageUrl,
		}
		res.Products = append(res.Products, productRes)
	}
	log.Println("gRPC - GetProducts result length:", len(res.Products))

	return res, nil
}

func (s *ShopProductServer) ProductExists(_ context.Context, payload *pb.CheckProductRequest) (*pb.CheckProductReponse, error) {
	productRepository := repositories.ProductRepository{DB: s.AppState.DB}

	result, err := productRepository.VerifyIsProductExistsByID(uint(payload.ProductId))
	if err != nil {
		return nil, err
	}

	return &pb.CheckProductReponse{IsExists: result}, nil
}
