package order

import (
	"github.com/buemura/event-driven-commerce/order-svc/internal/domain/common"
	"github.com/buemura/event-driven-commerce/order-svc/internal/domain/product"
)

type CreateOrderIn struct {
	CustomerId    string
	ProductList   []*product.Product
	PaymentMethod string
}

type GetManyOrdersIn struct {
	Page  int
	Items int
}

type GetManyOrdersOut struct {
	OrderList []*Order
	Meta      *common.PaginationMeta
}
