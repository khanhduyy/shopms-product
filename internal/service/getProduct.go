package service

import (
	"context"

	"github.com/khanhduyy/shopms-product/internal/domain"
	pb "github.com/khanhduyy/shopms-product/rpc/product"
)

func (p *productServiceImpl) GetAllProduct(ctx context.Context, request *pb.GetProductsRequest) (*pb.GetProductsResponse, error) {
	result, err := p.ProductRepository.FindAll(toProductQuery(request))
	if err != nil {
		return nil, err
	}
	var items []*pb.Product
	for _, v := range result.Items {
		items = append(items, v.GetProto())
	}

	return &pb.GetProductsResponse{
		Products: items,
	}, nil
}

func toProductQuery(productRequest *pb.GetProductsRequest) *domain.ProductQuery {
	return &domain.ProductQuery{
		Limit:  productRequest.Limit,
		Offset: productRequest.Offset,
	}
}
