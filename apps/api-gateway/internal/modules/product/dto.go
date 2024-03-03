package product

import (
	"github.com/buemura/event-driven-commerce/api-gateway/internal/shared"
)

type GetManyProductsIn struct {
	Page  int
	Items int
}

type GetManyProductsOut struct {
	ProductList []*Product
	Meta        *shared.PaginationMeta
}
