package main

import (
	"sync"

	"github.com/buemura/event-driven-commerce/payment-svc/config"
	"github.com/buemura/event-driven-commerce/payment-svc/internal/infra/queue/rabbit"
)

func init() {
	config.LoadEnv()
}

func main() {

	queueList := []string{
		"payment.create",
		"payment.process",
		"order.update",
	}

	var wg sync.WaitGroup
	wg.Add(len(queueList))

	for _, q := range queueList {
		go func() {
			defer wg.Done()
			rabbit.Consume(&rabbit.ConsumeIn{
				Queue: q,
			})
		}()
	}

	wg.Wait()
}
