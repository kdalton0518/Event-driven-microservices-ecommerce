package contracts

import (
	"github.com/buemura/event-driven-commerce/order-svc/internal/domain/product"
)

type ProductService interface {
	UpdateProductQuantity(id, quantity int) (*product.Product, error)
}
