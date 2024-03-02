package main

import (
	"log"
	"net"

	"github.com/buemura/event-driven-commerce/product-svc/internal/infra/grpc/server"
	"github.com/buemura/event-driven-commerce/product-svc/internal/infra/grpc/server/controllers"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":50050")
	if err != nil {
		log.Fatalf("cannot create grpc listener: %s", err)
	}

	s := grpc.NewServer()
	server.RegisterProductServiceServer(s, &controllers.ProductController{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to server grpc: %s", err)
	}
}
