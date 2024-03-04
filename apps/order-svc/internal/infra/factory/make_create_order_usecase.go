package factory

import (
	"github.com/buemura/event-driven-commerce/order-svc/internal/application/usecases"
	"github.com/buemura/event-driven-commerce/order-svc/internal/infra/database"
	"github.com/buemura/event-driven-commerce/order-svc/internal/infra/grpc/client"
)

func MakeCreateOrderUsecase() *usecases.CreateOrderUsecase {
	repo := database.NewPgxOrderRepository(database.Conn)
	pService := client.NewProductServiceClient()
	usecase := usecases.NewCreateOrderUsecase(repo, pService)
	return usecase
}
