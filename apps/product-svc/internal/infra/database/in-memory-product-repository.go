package database

import "github.com/buemura/event-driven-commerce/product-svc/internal/domain/product"

type InMemoryProductRepo struct {
	prod []*product.Product
}

func NewInMemoryProductRepo(prod []*product.Product) *InMemoryProductRepo {
	return &InMemoryProductRepo{
		prod: prod,
	}
}

func (r *InMemoryProductRepo) FindMany(in *product.GetManyProductsIn) (*product.ProductRepositoryPaginatedOut, error) {
	return &product.ProductRepositoryPaginatedOut{
		ProductList: r.prod,
		TotalCount:  len(r.prod),
	}, nil
}

func (r *InMemoryProductRepo) FindById(id int) (*product.Product, error) {
	var p *product.Product
	for _, v := range r.prod {
		if v.ID == id {
			p = v
			break
		}
	}
	return p, nil
}

func (r *InMemoryProductRepo) Update(newP *product.Product) (*product.Product, error) {
	var p *product.Product
	for _, v := range r.prod {
		if v.ID == newP.ID {
			p = v
			v = newP
			break
		}
	}

	if p == nil {
		r.prod = append(r.prod, p)
	}
	return newP, nil
}
