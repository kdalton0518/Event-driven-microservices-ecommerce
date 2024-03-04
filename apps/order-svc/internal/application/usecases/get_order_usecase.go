package usecases

import (
	"github.com/buemura/event-driven-commerce/order-svc/internal/domain/order"
)

type GetOrderUsecase struct {
	repo order.OrderRepository
}

func NewGetOrderUsecase(repo order.OrderRepository) *GetOrderUsecase {
	return &GetOrderUsecase{
		repo: repo,
	}
}

func (s *GetOrderUsecase) Execute(id string) (*order.Order, error) {
	ord, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	if ord == nil {
		return nil, order.ErrOrderNotFound
	}
	return ord, nil
}
