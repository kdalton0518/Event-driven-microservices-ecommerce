package presenter

import (
	"time"

	"github.com/buemura/event-driven-commerce/api-gateway/internal/modules/order"
	"github.com/buemura/event-driven-commerce/api-gateway/internal/modules/product"
	"github.com/buemura/event-driven-commerce/packages/pb"
)

func GrpcToDomainOrder(in *pb.OrderResponse) *order.Order {
	var productIdList []int
	for _, id := range in.ProductIdList {
		productIdList = append(productIdList, int(id))
	}

	createdAt, _ := time.Parse("2006-01-02T15:04:05.000Z", in.CreatedAt)
	updatedAt, _ := time.Parse("2006-01-02T15:04:05.000Z", in.UpdatedAt)

	return &order.Order{
		ID:            in.Id,
		CustomerId:    in.CustomerId,
		ProductIdList: productIdList,
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
