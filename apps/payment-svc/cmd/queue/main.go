package main

import (
	"sync"

	"github.com/buemura/event-driven-commerce/payment-svc/config"
	"github.com/buemura/event-driven-commerce/payment-svc/internal/infra/database"
	"github.com/buemura/event-driven-commerce/payment-svc/internal/infra/queue/rabbit"
)

func init() {
	config.LoadEnv()
	database.Connect()
	rabbit.DeclareQueueList()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(len(rabbit.QueueConsumerList))

	for _, q := range rabbit.QueueConsumerList {
		go func() {
			defer wg.Done()
			rabbit.Consume(&rabbit.ConsumeIn{
				Queue: q,
			})
		}()
	}

	wg.Wait()
}
