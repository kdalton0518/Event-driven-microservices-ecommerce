package factory

import (
	"github.com/buemura/event-driven-commerce/api-gateway/internal/infra/grpc/client"
	"github.com/buemura/event-driven-commerce/api-gateway/internal/modules/product"
)

func MakeProductDomainService() *product.ProductService {
	api := client.NewProductGrpcService()
	service := product.NewProductService(api)
	return service
}
