package payment

type CreatePaymentIn struct {
	OrderId       string        `json:"order_id"`
	PaymentMethod PaymentMethod `json:"payment_method"`
}

type CreatePaymentOut struct {
	OrderId string `json:"order_id"`
}

type ProcessPaymentIn struct {
	OrderId string `json:"order_id"`
}
