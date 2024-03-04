package main

import (
	"log"
	"net"

	"github.com/buemura/event-driven-commerce/order-svc/config"
	"github.com/buemura/event-driven-commerce/order-svc/internal/infra/database"
	"github.com/buemura/event-driven-commerce/order-svc/internal/infra/grpc/server/controllers"
	"github.com/buemura/event-driven-commerce/packages/pb"
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
	pb.RegisterOrderServiceServer(s, &controllers.OrderController{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to server grpc: %s", err)
	}
}
