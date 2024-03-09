package rabbit

import (
	"log"

	"github.com/buemura/event-driven-commerce/payment-svc/config"
	"github.com/buemura/event-driven-commerce/payment-svc/internal/infra/queue/controller"
	"github.com/buemura/event-driven-commerce/payment-svc/internal/infra/util"
	amqp "github.com/rabbitmq/amqp091-go"
)

type ConsumeIn struct {
	Queue string
}

func Consume(in *ConsumeIn) {
	conn, err := amqp.Dial(config.BROKER_URL)
	util.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	util.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	msgs, err := ch.Consume(
		in.Queue, // queue
		"",       // consumer
		true,     // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
	)
	util.FailOnError(err, "Failed to register a consumer")

	var forever chan struct{}
	go func() {
		for d := range msgs {
			log.Printf("\n")
			handleMessage(d)
		}
	}()

	log.Printf("RabbitMQ Consumer running for: Queue = %s", in.Queue)
	<-forever
}

func handleMessage(d amqp.Delivery) {

	switch d.RoutingKey {
	case "order.create":
		controller.CreateOrder(string(d.Body))
	case "order.update":
		controller.UpdateOrder(string(d.Body))
	case "payment.create":
		controller.CreatePayment(string(d.Body))
	case "payment.process":
		controller.ProcessPayment(string(d.Body))
	}
}
