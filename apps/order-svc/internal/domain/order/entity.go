package order

import (
	"time"

	"github.com/buemura/event-driven-commerce/order-svc/internal/domain/product"
	"github.com/google/uuid"
)

var WaitingPaymentStatus = "WAITING_PAYMENT"

type Order struct {
	ID            string
	CustomerId    string
	ProductList   []*product.Product
	TotalPrice    int
	Status        string
	PaymentMethod string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewOrder(in *CreateOrderIn) (*Order, error) {
	err := validate(in)
	if err != nil {
		return nil, err
	}

	totalPrice := calculateTotalPrice(in)
	if totalPrice == 0 {
		return nil, ErrOrderMissingProduct
	}

	return &Order{
		ID:            uuid.NewString(),
		CustomerId:    in.CustomerId,
		ProductList:   in.ProductList,
		TotalPrice:    totalPrice,
		Status:        WaitingPaymentStatus,
		PaymentMethod: in.PaymentMethod,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}, nil
}

func validate(in *CreateOrderIn) error {
	if len(in.ProductList) == 0 {
		return ErrOrderMissingProduct
	}
	if in.CustomerId == "" {
		return ErrOrderMissingCustomer
	}
	if in.PaymentMethod == "" {
		return ErrOrderMissingPaymentMethod
	}
	return nil
}

func calculateTotalPrice(in *CreateOrderIn) int {
	var total int
	for _, p := range in.ProductList {
		total += (p.Price * p.Quantity)
	}
	return total
}
