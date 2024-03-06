package order

import (
	"time"
)

type Order struct {
	ID            string          `json:"id"`
	CustomerId    string          `json:"customer_id"`
	ProductList   []*OrderProduct `json:"product_list"`
	TotalPrice    int             `json:"total_price"`
	Status        string          `json:"status"`
	PaymentMethod string          `json:"payment_method"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
}
