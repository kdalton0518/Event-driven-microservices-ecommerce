package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/buemura/event-driven-commerce/packages/pb"
	"github.com/buemura/event-driven-commerce/product-svc/config"
	"github.com/buemura/event-driven-commerce/product-svc/internal/infra/database"
	"github.com/buemura/event-driven-commerce/product-svc/internal/infra/grpc/controllers"
	"google.golang.org/grpc"
)

func init() {
	config.LoadEnv()
	database.Connect()
}

func main() {
	port := ":" + config.GRPC_PORT
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("cannot create grpc listener: %s", err)
	}

	s := grpc.NewServer()
	pb.RegisterProductServiceServer(s, &controllers.ProductController{})

	go func() {
		if err := s.Serve(listener); err != nil {
			log.Fatalf("failed to server grpc: %s", err)
		}
	}()

	log.Println("gRPC Server running at", port, "...")
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, os.Interrupt, syscall.SIGINT)
	<-stop

	log.Println("Stopping gRPC Server...")
	s.GracefulStop()
	log.Println("gRPC Server stopped")
}
