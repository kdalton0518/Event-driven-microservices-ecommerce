package presenter

import (
	"github.com/buemura/event-driven-commerce/order-svc/internal/domain/order"
	"github.com/buemura/event-driven-commerce/packages/pb"
)

func OrderToGrpc(in *order.Order) *pb.OrderResponse {
	var productIdList []int32
	for _, id := range in.ProductIdList {
		productIdList = append(productIdList, int32(id))
	}

	return &pb.OrderResponse{
		Id:            in.ID,
		CustomerId:    in.CustomerId,
		ProductIdList: productIdList,
		TotalPrice:    int64(in.TotalPrice),
		Status:        in.Status,
		PaymentMethod: in.PaymentMethod,
		CreatedAt:     in.CreatedAt.String(),
		UpdatedAt:     in.UpdatedAt.String(),
	}
}
