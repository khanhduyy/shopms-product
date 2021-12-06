package repository

import (
	"github.com/khanhduyy/shopms-common/client/db"
	q "github.com/khanhduyy/shopms-common/mvc/query"
	"github.com/khanhduyy/shopms-product/internal/domain"
)

type ProductRepository interface {
	FindAll(query *domain.ProductQuery) (*domain.ProductPage, error)
	Create(products []*domain.Product) ([]*domain.Product, error)
}

func NewProductRepository(db *db.Client) ProductRepository {
	return &productRepositoryIpml{
		tx: db,
	}
}

type productRepositoryIpml struct {
	tx *db.Client
}

func (p *productRepositoryIpml) FindAll(params *domain.ProductQuery) (*domain.ProductPage, error) {
	var query = q.NewQuery(p.tx).
		Select([]string{"p.*, p.Id"}).
		From(domain.ProductTable, "p").
		Page(params.Offset, params.Limit)
	var products []*domain.Product
	total, err := query.FindPage(&products)
	if err != nil {
		return nil, err
	}
	return &domain.ProductPage{
		Total: total,
		Items: products,
	}, nil
}

func (p *productRepositoryIpml) Create(products []*domain.Product) ([]*domain.Product, error) {
	var results = p.tx.Save(products)
	if results.Error != nil {
		return nil, results.Error
	}
	return products, nil
}
