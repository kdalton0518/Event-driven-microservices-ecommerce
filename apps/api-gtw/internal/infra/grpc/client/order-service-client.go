package client

import (
	"context"
	"log"

	"github.com/buemura/event-driven-commerce/api-gtw/config"
	"github.com/buemura/event-driven-commerce/api-gtw/internal/infra/presenter"
	"github.com/buemura/event-driven-commerce/api-gtw/internal/modules/order"
	"github.com/buemura/event-driven-commerce/api-gtw/internal/shared"
	"github.com/buemura/event-driven-commerce/packages/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OrderGrpcService struct{}

func NewOrderGrpcService() *OrderGrpcService {
	return &OrderGrpcService{}
}

func (*OrderGrpcService) GetOrder(id string) (*order.Order, error) {
	conn, err := grpc.Dial(config.GRPC_HOST_ORDER_SVC, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewOrderServiceClient(conn)

	log.Println("[GrpcClient][GetOrder] - Request order for id:", id)

	request := &pb.GetOrderRequest{Id: id}
	ord, err := client.GetOrder(context.Background(), request)
	if err != nil {
		log.Println("[GrpcClient][GetOrder] - Error:", err)
		return nil, err
	}

	return presenter.GrpcToDomainOrder(ord), nil
}

func (*OrderGrpcService) GetManyOrders(in *order.GetManyOrdersIn) (*order.GetManyOrdersOut, error) {
	conn, err := grpc.Dial(config.GRPC_HOST_ORDER_SVC, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	client := pb.NewOrderServiceClient(conn)
	log.Println("[GrpcClient][GetManyOrders] - Request product list")

	request := &pb.GetManyOrdersRequest{Page: int32(in.Page), Items: int32(in.Items)}
	res, err := client.GetManyOrders(context.Background(), request)
	if err != nil {
		log.Println("[GrpcClient][GetManyOrders] - Error:", err)
		return nil, err
	}

	var orderList []*order.Order
	for _, ord := range res.OrderList {
		orderList = append(orderList, presenter.GrpcToDomainOrder(ord))
	}

	return &order.GetManyOrdersOut{
		OrderList: orderList,
		Meta: &shared.PaginationMeta{
			Page:       int(res.Meta.Page),
			Items:      int(res.Meta.Items),
			TotalPages: int(res.Meta.TotalPages),
			TotalItems: int(res.Meta.TotalItems),
		},
	}, nil
}

func (*OrderGrpcService) CreateOrder(in *order.CreateOrderIn) (*order.Order, error) {
	conn, err := grpc.Dial(config.GRPC_HOST_ORDER_SVC, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial order-svc server: %v", err)
	}
	defer conn.Close()

	client := pb.NewOrderServiceClient(conn)
	log.Println("[GrpcClient][CreateOrder] - Request created order for customer:", in.CustomerId)

	var producList []*pb.CreateOrderRequest_OrderProduct
	for _, p := range in.ProductList {
		producList = append(producList, presenter.DomainOrderProductToGRPC(p))
	}

	request := &pb.CreateOrderRequest{
		CustomerId:    in.CustomerId,
		ProductList:   producList,
		PaymentMethod: in.PaymentMethod,
	}
	res, err := client.CreateOrder(context.Background(), request)
	if err != nil {
		log.Println("[GrpcClient][CreateOrder] - Error:", err)
		return nil, err
	}

	return presenter.GrpcToDomainOrder(res), nil
}
