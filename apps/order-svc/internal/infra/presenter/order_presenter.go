package presenter

import (
	"github.com/buemura/event-driven-commerce/order-svc/internal/domain/order"
	"github.com/buemura/event-driven-commerce/packages/pb"
)

func OrderToGrpc(in *order.Order) *pb.OrderResponse {
	var productList []*pb.OrderResponse_OrderProduct
	for _, p := range in.ProductList {
		productList = append(productList, &pb.OrderResponse_OrderProduct{
			Id:       int32(p.ID),
			Price:    int32(p.Price),
			Quantity: int32(p.Quantity),
		})
	}

	return &pb.OrderResponse{
		Id:            in.ID,
		CustomerId:    in.CustomerId,
		ProductList:   productList,
		TotalPrice:    int64(in.TotalPrice),
		Status:        in.Status,
		PaymentMethod: in.PaymentMethod,
		CreatedAt:     in.CreatedAt.Format("2006-01-02T15:04:05.000Z"),
		UpdatedAt:     in.UpdatedAt.Format("2006-01-02T15:04:05.000Z"),
	}
}
