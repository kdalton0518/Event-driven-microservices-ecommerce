package usecases

import (
	"testing"
	"time"

	"github.com/buemura/event-driven-commerce/order-svc/internal/domain/order"
	"github.com/buemura/event-driven-commerce/order-svc/internal/domain/product"
	"github.com/buemura/event-driven-commerce/order-svc/internal/infra/database"
	"github.com/stretchr/testify/assert"
)

func TestGetOrderUsecase(t *testing.T) {
	currentTime := time.Now()

	o := []*order.Order{
		{
			ID:            "order_1",
			CustomerId:    "customer_1",
			TotalPrice:    100,
			Status:        "WAITING_PAYMENT",
			PaymentMethod: "pix",
			CreatedAt:     currentTime,
			UpdatedAt:     currentTime,
			ProductList: []*product.Product{
				{ID: 1, Price: 100, Quantity: 1},
			},
		},
		{
			ID:            "order_2",
			CustomerId:    "customer_2",
			TotalPrice:    300,
			Status:        "WAITING_PAYMENT",
			PaymentMethod: "pix",
			CreatedAt:     currentTime,
			UpdatedAt:     currentTime,
			ProductList: []*product.Product{
				{ID: 1, Price: 100, Quantity: 1},
				{ID: 2, Price: 100, Quantity: 1},
				{ID: 3, Price: 100, Quantity: 1},
			},
		},
		{
			ID:            "order_3",
			CustomerId:    "customer_3",
			TotalPrice:    200,
			Status:        "SHIPPED",
			PaymentMethod: "pix",
			CreatedAt:     currentTime,
			UpdatedAt:     currentTime,
			ProductList: []*product.Product{
				{ID: 1, Price: 100, Quantity: 1},
				{ID: 2, Price: 100, Quantity: 1},
			},
		},
	}

	expected := *&order.Order{
		ID:            "order_3",
		CustomerId:    "customer_3",
		TotalPrice:    200,
		Status:        "SHIPPED",
		PaymentMethod: "pix",
		CreatedAt:     currentTime,
		UpdatedAt:     currentTime,
		ProductList: []*product.Product{
			{ID: 1, Price: 100, Quantity: 1},
			{ID: 2, Price: 100, Quantity: 1},
		},
	}

	repo := database.NewInMemoryOrderRepo(o)
	service := NewGetOrderUsecase(repo)

	t.Run("Return order struct", func(t *testing.T) {
		res, _ := service.Execute("order_3")
		assert.Equal(t, &expected, res)
	})
}
