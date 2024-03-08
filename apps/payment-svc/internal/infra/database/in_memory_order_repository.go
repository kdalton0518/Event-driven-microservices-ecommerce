package database

import "github.com/buemura/event-driven-commerce/payment-svc/internal/domain/order"

type InMemoryOrderRepo struct {
	order []*order.Order
}

func NewInMemoryOrderRepo(order []*order.Order) *InMemoryOrderRepo {
	return &InMemoryOrderRepo{
		order: order,
	}
}

func (r *InMemoryOrderRepo) FindMany(in *order.GetManyOrdersIn) (*order.OrderRepositoryPaginatedOut, error) {
	return &order.OrderRepositoryPaginatedOut{
		OrderList:  r.order,
		TotalCount: len(r.order),
	}, nil
}

func (r *InMemoryOrderRepo) FindById(id string) (*order.Order, error) {
	var o *order.Order
	for _, v := range r.order {
		if v.ID == id {
			o = v
			break
		}
	}
	return o, nil
}

func (r *InMemoryOrderRepo) Save(o *order.Order) (*order.Order, error) {
	r.order = append(r.order, o)
	return o, nil
}
