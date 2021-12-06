package service

import (
	"context"

	repo "github.com/khanhduyy/shopms-product/internal/repository"
	pb "github.com/khanhduyy/shopms-product/rpc/product"
)

type ProductService interface {
	GetAllProduct(ctx context.Context, empty *pb.GetProductsRequest) (*pb.GetProductsResponse, error)
	Save(ctx context.Context, productRequests []*pb.Product) (*pb.CreateProductResponse, error)
}

type productServiceImpl struct {
	ProductRepository repo.ProductRepository
}

func NewProductService(repository repo.ProductRepository) ProductService {
	return &productServiceImpl{
		ProductRepository: repository,
	}
}
