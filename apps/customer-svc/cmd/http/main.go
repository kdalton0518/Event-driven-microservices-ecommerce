package main

import (
	"context"
	"customer-svc/config"
	"customer-svc/internal/infra/database"
	"customer-svc/internal/infra/http/router"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	config.LoadEnv()
	database.Connect()
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

	fmt.Printf("Server running at %s ...", port)
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
