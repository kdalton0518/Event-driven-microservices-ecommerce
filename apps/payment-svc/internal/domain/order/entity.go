package order

import (
	"time"

	"github.com/buemura/event-driven-commerce/payment-svc/internal/domain/payment"
)

type Order struct {
	ID               string
	Amount           int
	Status           OrderStatus
	PaymentMethod    string
	PaymentOrderList []*payment.Payment
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func NewOrder(in *CreateOrderIn) (*Order, error) {

	err := validate(in)
	if err != nil {
		return nil, err
	}

	return &Order{
		ID:            in.OrderId,
		Amount:        in.Amount,
		Status:        StatusPending,
		PaymentMethod: in.PaymentMethod,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}, nil
}

func validate(in *CreateOrderIn) error {
	if in.OrderId == "" || in.Amount <= 0 || in.PaymentMethod == "" {
		return ErrInvalidArgument
	}
	if in.PaymentMethod != string(payment.CreditCard) && in.PaymentMethod != string(payment.Pix) {
		return ErrInvalidPaymentMethod
	}
	return nil
}
