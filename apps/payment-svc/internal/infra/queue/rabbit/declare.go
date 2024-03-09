package rabbit

import (
	"github.com/buemura/event-driven-commerce/payment-svc/config"
	amqp "github.com/rabbitmq/amqp091-go"
)

func DeclareQueueList() {
	conn, err := amqp.Dial(config.BROKER_URL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	for _, q := range QueueDeclareList {
		_, err := ch.QueueDeclare(
			q,     // name
			true,  // durable
			false, // delete when unused
			false, // exclusive
			false, // no-wait
			nil,   // arguments
		)
		failOnError(err, "Failed to declare a queue")
	}
}
