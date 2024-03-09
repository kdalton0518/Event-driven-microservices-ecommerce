package controller

import (
	"encoding/json"
	"log"

	"github.com/buemura/event-driven-commerce/payment-svc/internal/application/usecase"
	"github.com/buemura/event-driven-commerce/payment-svc/internal/domain/order"
	"github.com/buemura/event-driven-commerce/payment-svc/internal/infra/database"
	"github.com/buemura/event-driven-commerce/payment-svc/internal/infra/queue"
)

func CreateOrder(payload string) {
	var in *order.CreateOrderIn
	err := json.Unmarshal([]byte(payload), &in)
	if err != nil {
		log.Fatalf(err.Error())
	}

	repo := database.NewPgxOrderRepository()
	uc := usecase.NewOrderCreateUsecase(repo)
	o, err := uc.Execute(in)
	if err != nil {
		// queue.Publish(&queue.PublishIn{
		// 	Queue:   "order.create.dlq",
		// 	Payload: payload,
		// })
	}

	paymentCreate, _ := json.Marshal(&order.CreateOrderOut{
		OrderID:       o.ID,
		Amount:        o.Amount,
		PaymentMethod: o.PaymentMethod,
	})
	queue.Publish(&queue.PublishIn{
		RountingKey: "payment.create",
		Payload:     string(paymentCreate),
	})
}

func UpdateOrder(payload string) {
	var in *order.UpdateOrderIn
	err := json.Unmarshal([]byte(payload), &in)
	if err != nil {
		log.Fatalf(err.Error())
	}

	repo := database.NewPgxOrderRepository()
	uc := usecase.NewOrderUpdateUsecase(repo)
	o, err := uc.Execute(in)
	if err != nil {
		// queue.Publish(&queue.PublishIn{
		// 	Queue:   "order.create.dlq",
		// 	Payload: payload,
		// })
	}
	log.Println(o)
}
