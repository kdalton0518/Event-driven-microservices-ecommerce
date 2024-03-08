package order

import "errors"

var (
	ErrInvalidArgument      = errors.New("invalid argument")
	ErrInvalidPaymentMethod = errors.New("invalid payment method")
	ErrOrderNotFound        = errors.New("order not found")
)
