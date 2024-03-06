package factory

import (
	"github.com/buemura/event-driven-commerce/api-gtw/internal/infra/grpc/client"
	"github.com/buemura/event-driven-commerce/api-gtw/internal/modules/customer"
)

func MakeCustomerDomainService() *customer.CustomerService {
	api := client.NewCustomerGrpcService()
	service := customer.NewCustomerService(api)
	return service
}
