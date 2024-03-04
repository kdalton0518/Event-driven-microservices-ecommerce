package usecases

import "github.com/buemura/event-driven-commerce/order-svc/internal/domain/order"

type CreateOrderUsecase struct {
	repo order.OrderRepository
}

func NewCreateOrderUsecase(repo order.OrderRepository) *CreateOrderUsecase {
	return &CreateOrderUsecase{
		repo: repo,
	}
}

func (s *CreateOrderUsecase) Execute(in *order.CreateOrderIn) (*order.Order, error) {
	o, err := order.NewOrder(in)
	if err != nil {
		return nil, err
	}

	res, err := s.repo.Save(o)
	if err != nil {
		return nil, err
	}

	return res, nil
}
