package api

import (
	"context"

	log "github.com/khanhduyy/shopms-common/logger"
	pb "github.com/khanhduyy/shopms-product/rpc/product"
)

func (p *ProductApi) GetAllProduct(ctx context.Context, request *pb.GetProductsRequest) (*pb.GetProductsResponse, error) {
	result, err := p.ProductService.GetAllProduct(ctx, request)
	if err != nil {
		log.Errorf("Cound not get products on error %v", err)
		return nil, err
	}
	return result, nil
}
