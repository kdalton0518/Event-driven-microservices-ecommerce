package usecases

import (
	"github.com/buemura/event-driven-commerce/order-svc/internal/application/contracts"
	"github.com/buemura/event-driven-commerce/order-svc/internal/domain/order"
)

type CreateOrderUsecase struct {
	repo           order.OrderRepository
	productService contracts.ProductService
}

func NewCreateOrderUsecase(
	repo order.OrderRepository,
	productService contracts.ProductService,
) *CreateOrderUsecase {
	return &CreateOrderUsecase{
		repo:           repo,
		productService: productService,
	}
}

func (s *CreateOrderUsecase) Execute(in *order.CreateOrderIn) (*order.Order, error) {
	o, err := order.NewOrder(in)
	if err != nil {
		return nil, err
	}

	for _, p := range in.ProductList {
		_, err := s.productService.UpdateProductQuantity(p.ID, (p.Quantity * -1))
		if err != nil {
			return nil, err
		}
	}

	res, err := s.repo.Save(o)
	if err != nil {
		return nil, err
	}

	return res, nil
}
