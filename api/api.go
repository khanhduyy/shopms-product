package api

import "github.com/khanhduyy/shopms-product/internal/service"

type ProductApi struct {
	ProductService service.ProductService
}

func NewProductApi(productService service.ProductService) *ProductApi {
	return &ProductApi{
		ProductService: productService,
	}
}
