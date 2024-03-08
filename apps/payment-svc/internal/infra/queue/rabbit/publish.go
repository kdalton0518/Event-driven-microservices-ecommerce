package rabbit

import (
	"context"
	"log"
	"time"

	"github.com/buemura/event-driven-commerce/payment-svc/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

type PublishIn struct {
	Exchange    string
	Queue       string
	RountingKey string
	Payload     string
}

func Send(in *PublishIn) {
	conn, err := amqp.Dial(config.BROKER_URL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// q, err := ch.QueueDeclare(
	// 	queue, // name
	// 	false, // durable
	// 	false, // delete when unused
	// 	false, // exclusive
	// 	false, // no-wait
	// 	nil,   // arguments
	// )
	// failOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		in.Exchange,    // exchange
		in.RountingKey, // routing key
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(in.Payload),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", in.Payload)
}
