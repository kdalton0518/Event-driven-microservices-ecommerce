package presenter

import (
	"testing"
	"time"

	"github.com/buemura/event-driven-commerce/order-svc/internal/domain/order"
	"github.com/buemura/event-driven-commerce/order-svc/internal/domain/product"
	"github.com/buemura/event-driven-commerce/packages/pb"
	"github.com/stretchr/testify/assert"
)

func TestOrderPresenter(t *testing.T) {
	currentTime := time.Now()

	in := &order.Order{
		ID:         "any_order_id",
		CustomerId: "any_customer_id",
		ProductList: []*product.Product{
			{ID: 1, Price: 100, Quantity: 1},
		},
		TotalPrice:    100,
		Status:        "any_status",
		PaymentMethod: "any_payment_method",
		CreatedAt:     currentTime,
		UpdatedAt:     currentTime,
	}

	expected := &pb.OrderResponse{
		Id:         "any_order_id",
		CustomerId: "any_customer_id",
		ProductList: []*pb.OrderResponse_OrderProduct{
			{Id: 1, Price: 100, Quantity: 1},
		},
		TotalPrice:    100,
		Status:        "any_status",
		PaymentMethod: "any_payment_method",
		CreatedAt:     currentTime.Format("2006-01-02T15:04:05.000Z"),
		UpdatedAt:     currentTime.Format("2006-01-02T15:04:05.000Z"),
	}

	t.Run("DomainOrderToGrpc return the right structure", func(t *testing.T) {

		out := DomainOrderToGrpc(in)
		assert.Equal(t, out, expected)
	})
}
