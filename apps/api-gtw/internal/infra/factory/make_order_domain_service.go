package factory

import (
	"github.com/buemura/event-driven-commerce/api-gtw/internal/infra/grpc/client"
	"github.com/buemura/event-driven-commerce/api-gtw/internal/modules/order"
)

func MakeOrderDomainService() *order.OrderService {
	api := client.NewOrderGrpcService()
	service := order.NewOrderService(api)
	return service
}
