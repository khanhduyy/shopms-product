package repository

import (
	"github.com/khanhduyy/shopms-common/client/db"
	q "github.com/khanhduyy/shopms-common/mvc/query"
	"github.com/khanhduyy/shopms-product/internal/domain"
)

type ProductRepository interface {
	FindAll(query *domain.ProductQuery) (*domain.ProductPage, error)
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
		Select([]string{"p.*"}).
		From(domain.ProductTable, "p").
		Page(params.Offset, params.Limit)

}
