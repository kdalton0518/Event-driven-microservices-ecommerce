package factory

import (
	"github.com/buemura/event-driven-commerce/order-svc/internal/application/usecases"
	"github.com/buemura/event-driven-commerce/order-svc/internal/infra/database"
)

func MakeGetManyOrdersUsecase() *usecases.GetManyOrdersUsecase {
	repo := database.NewPgxOrderRepository(database.Conn)
	usecase := usecases.NewGetManyOrdersUsecase(repo)
	return usecase
}
