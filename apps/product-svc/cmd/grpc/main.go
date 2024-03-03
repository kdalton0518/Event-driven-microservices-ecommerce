package main

import (
	"log"
	"net"

	"github.com/buemura/event-driven-commerce/product-svc/config"
	"github.com/buemura/event-driven-commerce/product-svc/internal/infra/database"
	"github.com/buemura/event-driven-commerce/product-svc/internal/infra/grpc/server"
	"github.com/buemura/event-driven-commerce/product-svc/internal/infra/grpc/server/controllers"
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
	server.RegisterProductServiceServer(s, &controllers.ProductController{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to server grpc: %s", err)
	}
}
