package queue

import (
	"context"
	"log"
	"time"

	"github.com/buemura/event-driven-commerce/payment-svc/config"
	"github.com/buemura/event-driven-commerce/payment-svc/internal/infra/util"
	amqp "github.com/rabbitmq/amqp091-go"
)

type PublishIn struct {
	Exchange    string
	Queue       string
	RountingKey string
	Payload     string
}

func Publish(in *PublishIn) {
	conn, err := amqp.Dial(config.BROKER_URL)
	util.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	util.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	util.FailOnError(err, "Failed to declare a queue")

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
	util.FailOnError(err, "Failed to publish a message")
	log.Printf("[Queue][Publish] - Sent message to %s: \n", in.RountingKey)
}
