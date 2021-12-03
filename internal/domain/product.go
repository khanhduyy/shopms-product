package domain

import (
	e "github.com/khanhduyy/shopms-common/mvc/entity"
	pb "github.com/khanhduyy/shopms-product/rpc/product"
)

const (
	ProductTable = "Product"
)

type Product struct {
	*e.Base
	Name        string  `gorm:"column:Name"`
	Description string  `gorm:"column:Description;"`
	Quantity    int64   `gorm:"column:Quantity;"`
	Price       float64 `gorm:"column:Price;"`
}

func (Product) TableName() string {
	return ProductTable
}

func NewProduct(product *pb.Product) *Product {
	if product == nil {
		return nil
	}
	return &Product{
		Base:        e.NewBaseInt64(product.GetId()),
		Name:        product.GetName(),
		Description: product.GetDescription(),
		Quantity:    product.GetQuantity(),
		Price:       float64(product.GetPrice()),
	}
}

func (p *Product) GetProto() *pb.Product {
	if p == nil {
		return nil
	}
	return &pb.Product{
		Id:          int64(p.Base.Id),
		Name:        p.Name,
		Description: p.Description,
		Quantity:    p.Quantity,
		Price:       float32(p.Price),
	}
}

type ProductPage struct {
	Total int64
	Items []*Product
}

type ProductQuery struct {
	Sort, Order   string
	Offset, Limit int64
}
