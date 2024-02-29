package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"user-svc/config"
	"user-svc/internal/infra/http/router"
)

func init() {
	config.LoadEnv()
}

func main() {
	router.SetupRouters()
	server := &http.Server{Addr: ":9001"}

	go func() {
		if err := server.ListenAndServe(); err != nil && http.ErrServerClosed != err {
			panic(err)
		}
	}()

	fmt.Println("Server running at 9001...")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, os.Interrupt, syscall.SIGINT)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	fmt.Println("Stopping...")

	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	}
	fmt.Println("Server stopped")
}
