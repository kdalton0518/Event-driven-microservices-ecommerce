package main

import "github.com/buemura/event-driven-commerce/payment-svc/internal/infra/queue"

func main() {
	queue.Publish(&queue.PublishIn{
		RountingKey: "payment.create",
		Payload:     "hahahah",
	})
}
