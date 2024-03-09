package order

import "github.com/buemura/event-driven-commerce/payment-svc/internal/domain/common"

type CreateOrderIn struct {
	OrderId       string `json:"order_id"`
	Amount        int    `json:"amount"`
	PaymentMethod string `json:"payment_method"`
}

type CreateOrderOut struct {
	OrderID       string `json:"order_id"`
	Amount        int    `json:"amount"`
	PaymentMethod string `json:"payment_method"`
}

type UpdateOrderIn struct {
	OrderId string      `json:"order_id"`
	Status  OrderStatus `json:"status"`
}

type GetManyOrdersIn struct {
	Page  int
	Items int
}

type GetManyOrdersOut struct {
	OrderList []*Order
	Meta      *common.PaginationMeta
}
