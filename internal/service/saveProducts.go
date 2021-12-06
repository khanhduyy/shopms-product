package service

import (
	"context"

	"github.com/khanhduyy/shopms-product/internal/domain"
	pb "github.com/khanhduyy/shopms-product/rpc/product"
)

func (p *productServiceImpl) Save(ctx context.Context,
	productRequests []*pb.Product) (*pb.CreateProductResponse, error) {
	products, err := transform(productRequests)
	if err != nil {
		return nil, err
	}
	if _, err := p.ProductRepository.Create(products); err != nil {
		return nil, err
	}
	return newCreateProductsResponse(products), nil
}

func transform(productRequests []*pb.Product) (products []*domain.Product, err error) {
	for _, product := range productRequests {
		products = append(products, domain.NewProduct(product))
	}
	return products, nil
}

func newCreateProductsResponse(products []*domain.Product) *pb.CreateProductResponse {
	var ids []int64
	for _, v := range products {
		ids = append(ids, int64(v.Id))
	}
	return &pb.CreateProductResponse{
		ProductIds: ids,
	}
}
