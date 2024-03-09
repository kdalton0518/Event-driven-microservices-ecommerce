package queue

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type PublishIn struct {
	Exchange    string
	Queue       string
	RountingKey string
	Payload     string
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func Publish(in *PublishIn) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	// conn, err := amqp.Dial(config.BROKER_URL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	failOnError(err, "Failed to declare a queue")

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
	log.Printf("[Queue][Publish] - Sent message to %s: \n", in.RountingKey)
}
