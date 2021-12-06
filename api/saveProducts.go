package api

import (
	"context"

	log "github.com/khanhduyy/shopms-common/logger"
	pb "github.com/khanhduyy/shopms-product/rpc/product"
)

func (p *ProductApi) CreateProducts(ctx context.Context,
	productRequests *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	result, err := p.ProductService.Save(ctx, productRequests.Products)
	if err != nil {
		log.Errorf("Cound not save products on error %v", err)
		return nil, err
	}
	return result, nil
}
