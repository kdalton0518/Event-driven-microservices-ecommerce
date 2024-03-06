package order

import (
	"github.com/buemura/event-driven-commerce/api-gtw/internal/modules/product"
	"github.com/buemura/event-driven-commerce/api-gtw/internal/shared"
)

type CreateOrderIn struct {
	CustomerId    string             `json:"customer_id"`
	ProductList   []*product.Product `json:"product_list"`
	PaymentMethod string             `json:"payment_method"`
}

type GetManyOrdersIn struct {
	Page  int `json:"page"`
	Items int `json:"items"`
}

type GetManyOrdersOut struct {
	OrderList []*Order               `json:"order_list"`
	Meta      *shared.PaginationMeta `json:"meta"`
}

type OrderProduct struct {
	Id       int `json:"id"`
	Price    int `json:"price"`
	Quantity int `json:"quantity"`
}
