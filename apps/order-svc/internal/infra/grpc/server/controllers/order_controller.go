package controllers

import (
	"context"
	"log"

	"github.com/buemura/event-driven-commerce/order-svc/internal/domain/order"
	"github.com/buemura/event-driven-commerce/order-svc/internal/domain/product"
	"github.com/buemura/event-driven-commerce/order-svc/internal/infra/factory"
	"github.com/buemura/event-driven-commerce/order-svc/internal/infra/grpc/server/helper"
	"github.com/buemura/event-driven-commerce/order-svc/internal/presenter"
	"github.com/buemura/event-driven-commerce/packages/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OrderController struct {
	pb.UnimplementedOrderServiceServer
}

func (c OrderController) GetManyOrders(
	ctx context.Context,
	in *pb.GetManyOrdersRequest,
) (*pb.GetManyOrdersResponse, error) {
	log.Println("[GrpcServer][GetManyOrders] - Incoming request")
	var page, items int = 1, 10
	if in.Page != 0 {
		page = int(in.Page)
	}
	if in.Items != 0 {
		items = int(in.Items)
	}

	usecase := factory.MakeGetManyOrdersUsecase()

	res, err := usecase.Execute(&order.GetManyOrdersIn{
		Page:  page,
		Items: items,
	})
	if err != nil {
		log.Println("[GrpcServer][GetManyOrders] - Error:", err.Error())
		return nil, helper.HandleGrpcError(err)
	}

	var orderList []*pb.OrderResponse
	for _, o := range res.OrderList {
		ord := presenter.DomainOrderToGrpc(o)
		orderList = append(orderList, ord)
	}
	return &pb.GetManyOrdersResponse{
		OrderList: orderList,
		Meta: &pb.GetManyOrdersResponse_PaginationMeta{
			Page:       int32(res.Meta.Page),
			Items:      int32(res.Meta.Items),
			TotalPages: int32(res.Meta.TotalPages),
			TotalItems: int32(res.Meta.TotalItems),
		},
	}, nil
}

func (c OrderController) GetOrder(
	ctx context.Context,
	in *pb.GetOrderRequest,
) (*pb.OrderResponse, error) {
	log.Println("[GrpcServer][GetOrder] - Incoming request for id:", in.Id)
	if in.Id == "" {
		log.Println("[GrpcServer][GetOrder] - Error: missing id parameter")
		return nil, status.Error(codes.InvalidArgument, "missing id parameter")
	}

	usecase := factory.MakeGetOrderUsecase()

	o, err := usecase.Execute(in.Id)
	if err != nil {
		log.Println("[GrpcServer][GetOrder] - Error:", err.Error())
		return nil, helper.HandleGrpcError(err)
	}

	return presenter.DomainOrderToGrpc(o), nil
}

func (c OrderController) CreateOrder(
	ctx context.Context,
	in *pb.CreateOrderRequest,
) (*pb.OrderResponse, error) {
	log.Println("[GrpcServer][CreateOrder] - Incoming request for customer:", in.CustomerId)
	if in.CustomerId == "" {
		log.Println("[GrpcServer][CreateOrder] - Error: missing customer_id parameter")
		return nil, status.Error(codes.InvalidArgument, "missing customer_id parameter")
	}
	if in.PaymentMethod == "" {
		log.Println("[GrpcServer][CreateOrder] - Error: missing payment_method parameter")
		return nil, status.Error(codes.InvalidArgument, "missing payment_method parameter")
	}
	if len(in.ProductList) == 0 {
		log.Println("[GrpcServer][CreateOrder] - Error: missing product parameter")
		return nil, status.Error(codes.InvalidArgument, "missing product parameter")
	}

	usecase := factory.MakeCreateOrderUsecase()

	var productList []*product.Product
	for _, p := range in.ProductList {
		productList = append(productList, &product.Product{
			ID:       int(p.Id),
			Price:    int(p.Price),
			Quantity: int(p.Quantity),
		})
	}

	o, err := usecase.Execute(&order.CreateOrderIn{
		CustomerId:    in.CustomerId,
		ProductList:   productList,
		PaymentMethod: in.PaymentMethod,
	})
	if err != nil {
		log.Println("[GrpcServer][CreateOrder] - Error:", err.Error())
		return nil, helper.HandleGrpcError(err)
	}
	return presenter.DomainOrderToGrpc(o), nil
}
