package controller

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/buemura/event-driven-commerce/payment-svc/internal/application/usecase"
	"github.com/buemura/event-driven-commerce/payment-svc/internal/domain/payment"
	"github.com/buemura/event-driven-commerce/payment-svc/internal/infra/database"
	"github.com/buemura/event-driven-commerce/payment-svc/internal/infra/queue"
)

func CreatePayment(payload string) {
	var in *payment.CreatePaymentIn
	err := json.Unmarshal([]byte(payload), &in)
	if err != nil {
		log.Fatalf(err.Error())
	}

	repo := database.NewPgxPaymentRepository()
	uc := usecase.NewPaymentCreateUsecase(repo)

	p, err := uc.Execute(in)
	if err != nil {
		// queue.Publish(&queue.PublishIn{
		// 	Queue:   "payment.create.dlq",
		// 	Payload: payload,
		// })
	}

	processPaymentPayload, _ := json.Marshal(&payment.CreatePaymentOut{
		OrderId: p.OrderId,
	})
	queue.Publish(&queue.PublishIn{
		RountingKey: "payment.process",
		Payload:     string(processPaymentPayload),
	})
}

func ProcessPayment(payload string) {
	var in *payment.ProcessPaymentIn
	err := json.Unmarshal([]byte(payload), &in)
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println(in)
}
