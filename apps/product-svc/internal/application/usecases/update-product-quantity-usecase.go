package usecases

import (
	"github.com/buemura/event-driven-commerce/product-svc/internal/domain/product"
)

type UpdateProductQuantityUsecase struct {
	repo product.ProductRepository
}

func NewUpdateProductQuantityUsecase(repo product.ProductRepository) *UpdateProductQuantityUsecase {
	return &UpdateProductQuantityUsecase{
		repo: repo,
	}
}

func (s *UpdateProductQuantityUsecase) Execute(in *product.UpdateProductQuantityIn) (*product.Product, error) {
	prod, err := s.repo.FindById(in.ID)
	if err != nil {
		return nil, err
	}
	if prod == nil {
		return nil, product.ErrProductNotFound
	}

	prod.Quantity += in.Quantity

	if prod.Quantity < 0 {
		return nil, product.ErrProductInsufficientQuantity
	}

	res, err := s.repo.Update(prod)
	if err != nil {
		return nil, err
	}
	return res, nil
}
