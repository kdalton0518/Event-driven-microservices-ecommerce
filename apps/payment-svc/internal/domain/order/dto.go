package order

import "github.com/buemura/event-driven-commerce/payment-svc/internal/domain/common"

type CreateOrderIn struct {
	OrderId       string
	Amount        int
	PaymentMethod string
}

type UpdateOrderIn struct {
	OrderId string
	Status  OrderStatus
}

type GetManyOrdersIn struct {
	Page  int
	Items int
}

type GetManyOrdersOut struct {
	OrderList []*Order
	Meta      *common.PaginationMeta
}
