package order

import "github.com/buemura/event-driven-commerce/order-svc/internal/domain/common"

type CreateOrderIn struct {
	CustomerId    string
	ProductIdList []int
	TotalPrice    int
	Status        string
}

type GetManyOrdersIn struct {
	Page  int
	Items int
}

type GetManyProductsOut struct {
	OrderList []*Order
	Meta      *common.PaginationMeta
}
