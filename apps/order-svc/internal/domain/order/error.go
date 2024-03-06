package order

import "errors"

var (
	ErrOrderNotFound             = errors.New("order not found")
	ErrOrderMissingCustomer      = errors.New("order has no customer")
	ErrOrderMissingProduct       = errors.New("order has no product")
	ErrOrderMissingPaymentMethod = errors.New("order has no payment method")
)
