package order

import (
	"time"

	"github.com/google/uuid"
)

var WaitingPaymentStatus = "WAITING_PAYMENT"

type Order struct {
	ID            string
	CustomerId    string
	ProductIdList []int
	TotalPrice    int
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewOrder(in *CreateOrderIn) (*Order, error) {
	return &Order{
		ID:            uuid.NewString(),
		CustomerId:    in.CustomerId,
		ProductIdList: in.ProductIdList,
		TotalPrice:    in.TotalPrice,
		Status:        WaitingPaymentStatus,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}, nil
}
