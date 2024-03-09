package rabbit

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
