package product

import "github.com/buemura/event-driven-commerce/product-svc/internal/domain/common"

type GetManyProductsIn struct {
	Page  int
	Items int
}

type GetManyProductsOut struct {
	ProductList []*Product
	Meta        *common.PaginationMeta
}

type UpdateProductQuantityIn struct {
	ID       int
	Quantity int
}
