package payment

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID        string
	OrderId   string
	Status    PaymentStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewPayment(in *CreatePaymentIn) (*Payment, error) {
	if in.OrderId == "" {
		return nil, ErrCannotCreateWithoutOrderId
	}

	return &Payment{
		ID:        uuid.NewString(),
		OrderId:   in.OrderId,
		Status:    PaymentPending,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
