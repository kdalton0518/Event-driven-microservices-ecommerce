package order

import (
	"time"
)

type Order struct {
	ID            string    `json:"id"`
	CustomerId    string    `json:"customer_id"`
	ProductIdList []int     `json:"product_id_list"`
	TotalPrice    int       `json:"total_price"`
	Status        string    `json:"status"`
	PaymentMethod string    `json:"payment_method"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
