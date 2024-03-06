package factory

import (
	"github.com/buemura/event-driven-commerce/api-gtw/internal/infra/grpc/client"
	"github.com/buemura/event-driven-commerce/api-gtw/internal/modules/product"
)

func MakeProductDomainService() *product.ProductService {
	api := client.NewProductGrpcService()
	service := product.NewProductService(api)
	return service
}
