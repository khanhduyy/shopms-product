package service

import (
	"context"

	repo "github.com/khanhduyy/shopms-product/internal/repository"
	pb "github.com/khanhduyy/shopms-product/rpc/product"
)

type ProductService interface {
	GetAllProduct(ctx context.Context, empty *pb.Empty) (*pb.GetProductsResponse, error)
}

type productServiceImpl struct {
	ProductRepository repo.ProductRepository
}

func NewProductService(repository repo.ProductRepository) ProductService {
	return &productServiceImpl{
		ProductRepository: repository,
	}
}
