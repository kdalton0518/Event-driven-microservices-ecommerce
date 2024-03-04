package usecases

import "github.com/buemura/event-driven-commerce/product-svc/internal/domain/product"

type GetProductUsecase struct {
	repo product.ProductRepository
}

func NewGetProductUsecase(repo product.ProductRepository) *GetProductUsecase {
	return &GetProductUsecase{
		repo: repo,
	}
}

func (s *GetProductUsecase) Execute(id int) (*product.Product, error) {
	prod, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	if prod == nil {
		return nil, product.ErrProductNotFound
	}
	return prod, nil
}
