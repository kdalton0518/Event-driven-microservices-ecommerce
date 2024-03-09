package rabbit

import (
	"github.com/buemura/event-driven-commerce/payment-svc/config"
	"github.com/buemura/event-driven-commerce/payment-svc/internal/infra/util"
	amqp "github.com/rabbitmq/amqp091-go"
)

var QueueDeclareList []string = []string{
	"payment.create",
	"payment.create.dlq",
	"payment.process",
	"payment.process.dlq",
	"order.create",
	"order.create.dlq",
	"order.update",
	"order.update.dlq",
}

var QueueConsumerList []string = []string{
	"payment.create",
	"payment.process",
	"order.create",
	"order.update",
}

func DeclareQueueList() {
	conn, err := amqp.Dial(config.BROKER_URL)
	util.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	util.FailOnError(err, "Failed to open a channel")
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
		util.FailOnError(err, "Failed to declare a queue")
	}
}
