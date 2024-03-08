package main

import (
	"github.com/buemura/event-driven-commerce/payment-svc/config"
	"github.com/buemura/event-driven-commerce/payment-svc/internal/infra/queue/rabbit"
)

func init() {
	config.LoadEnv()
}

func main() {
	rabbit.Consume(&rabbit.ConsumeIn{
		Queue: "payment.create",
	})
}
