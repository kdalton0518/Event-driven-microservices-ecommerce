package usecase

import (
	"github.com/buemura/event-driven-commerce/payment-svc/internal/domain/order"
)

type OrderUpdateUsecase struct {
	repo order.OrderRepository
}

func NewOrderUpdateUsecase(repo order.OrderRepository) *OrderUpdateUsecase {
	return &OrderUpdateUsecase{
		repo: repo,
	}
}

func (u *OrderUpdateUsecase) Execute(in *order.UpdateOrderIn) (*order.Order, error) {
	o, err := u.repo.FindById(in.OrderId)
	if err != nil {
		return nil, order.ErrOrderNotFound
	}

	o.Status = in.Status

	err = u.repo.Update(o.ID, string(o.Status))
	if err != nil {
		return nil, err
	}

	return o, nil
}
