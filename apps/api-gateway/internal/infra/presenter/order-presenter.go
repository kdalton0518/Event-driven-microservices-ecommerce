package presenter

import (
	"time"

	"github.com/buemura/event-driven-commerce/api-gateway/internal/modules/order"
	"github.com/buemura/event-driven-commerce/api-gateway/internal/modules/product"
	"github.com/buemura/event-driven-commerce/packages/pb"
)

func GrpcToDomainOrder(in *pb.OrderResponse) *order.Order {
	var productList []*order.OrderProduct
	for _, p := range in.ProductList {
		productList = append(productList, &order.OrderProduct{
			Id:       int(p.Id),
			Price:    int(p.Price),
			Quantity: int(p.Quantity),
		})
	}

	createdAt, _ := time.Parse("2006-01-02T15:04:05.000Z", in.CreatedAt)
	updatedAt, _ := time.Parse("2006-01-02T15:04:05.000Z", in.UpdatedAt)

	return &order.Order{
		ID:            in.Id,
		CustomerId:    in.CustomerId,
		ProductList:   productList,
		TotalPrice:    int(in.TotalPrice),
		Status:        in.Status,
		PaymentMethod: in.PaymentMethod,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}
}

func DomainOrderProductToGRPC(in *product.Product) *pb.CreateOrderRequest_OrderProduct {
	return &pb.CreateOrderRequest_OrderProduct{
		Id:       int32(in.ID),
		Name:     in.Name,
		Price:    int32(in.Price),
		Quantity: int32(in.Quantity),
		ImageUrl: in.ImageUrl,
	}
}
