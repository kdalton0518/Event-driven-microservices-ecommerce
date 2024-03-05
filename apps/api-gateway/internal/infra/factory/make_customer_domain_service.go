package factory

import (
	"github.com/buemura/event-driven-commerce/api-gateway/internal/infra/grpc/client"
	"github.com/buemura/event-driven-commerce/api-gateway/internal/modules/customer"
)

func MakeCustomerDomainService() *customer.CustomerService {
	api := client.NewCustomerGrpcService()
	service := customer.NewCustomerService(api)
	return service
}
