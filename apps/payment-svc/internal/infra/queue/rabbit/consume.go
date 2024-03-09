package rabbit

import (
	"log"

	"github.com/buemura/event-driven-commerce/payment-svc/config"
	"github.com/buemura/event-driven-commerce/payment-svc/internal/infra/queue/controller"
	amqp "github.com/rabbitmq/amqp091-go"
)

type ConsumeIn struct {
	Queue string
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func Consume(in *ConsumeIn) {
	conn, err := amqp.Dial(config.BROKER_URL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		in.Queue, // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)
	failOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	var forever chan struct{}
	go func() {
		for d := range msgs {
			log.Printf("\n")
			handleMessage(d)
		}
	}()

	log.Printf("RabbitMQ Consumer running for: Queue = %s", q.Name)
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
