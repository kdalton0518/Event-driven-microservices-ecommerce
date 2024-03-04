package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/buemura/event-driven-commerce/api-gateway/config"
	"github.com/buemura/event-driven-commerce/api-gateway/internal/infra/http/router"
)

func init() {
	config.LoadEnv()
}

func main() {
	router.SetupRouters()
	port := ":" + config.PORT
	server := &http.Server{Addr: port}

	go func() {
		if err := server.ListenAndServe(); err != nil && http.ErrServerClosed != err {
			panic(err)
		}
	}()

	log.Println("HTTP Server running at", port, "...")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, os.Interrupt, syscall.SIGINT)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	log.Println("Stopping HTTP Server...")

	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	}
	log.Println("HTTP Server stopped")
}
