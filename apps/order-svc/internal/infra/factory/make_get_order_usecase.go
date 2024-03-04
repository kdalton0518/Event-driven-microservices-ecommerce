package factory

import (
	"github.com/buemura/event-driven-commerce/order-svc/internal/application/usecases"
	"github.com/buemura/event-driven-commerce/order-svc/internal/infra/database"
)

func MakeGetOrderUsecase() *usecases.GetOrderUsecase {
	repo := database.NewPgxOrderRepository(database.Conn)
	usecase := usecases.NewGetOrderUsecase(repo)
	return usecase
}
