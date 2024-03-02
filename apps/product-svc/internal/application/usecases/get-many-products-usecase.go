package usecases

import "github.com/buemura/event-driven-commerce/product-svc/internal/domain/product"

type GetManyProductUsecase struct {
	repo product.ProductRepository
}

func NewGetManyProductUsecase(repo product.ProductRepository) *GetManyProductUsecase {
	return &GetManyProductUsecase{
		repo: repo,
	}
}

func (uc *GetManyProductUsecase) Execute() ([]*product.Product, error) {
	prod, err := uc.repo.FindMany()
	if err != nil {
		return nil, err
	}
	return prod, nil
}
