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
	PaymentMethod string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func NewOrder(in *CreateOrderIn) (*Order, error) {
	err := validate(in)
	if err != nil {
		return nil, err
	}

	ids, totalPrice := parseProductAndCalculateTotalPrice(in)
	if totalPrice == 0 {
		return nil, ErrOrderMissingProduct
	}

	return &Order{
		ID:            uuid.NewString(),
		CustomerId:    in.CustomerId,
		ProductIdList: ids,
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

func parseProductAndCalculateTotalPrice(in *CreateOrderIn) ([]int, int) {
	var idList []int
	var total int
	for _, p := range in.ProductList {
		idList = append(idList, p.ID)
		total += (p.Price * p.Quantity)
	}
	return idList, total
}
