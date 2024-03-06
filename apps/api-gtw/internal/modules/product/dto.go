package product

import (
	"github.com/buemura/event-driven-commerce/api-gtw/internal/shared"
)

type GetManyProductsIn struct {
	Page  int `json:"page"`
	Items int `json:"items"`
}

type GetManyProductsOut struct {
	ProductList []*Product             `json:"product_list"`
	Meta        *shared.PaginationMeta `json:"meta"`
}
