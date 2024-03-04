package order

import "errors"

var ErrOrderNotFound = errors.New("order not found")
var ErrOrderMissingCustomer = errors.New("order has no customer")
var ErrOrderMissingProduct = errors.New("order has no product")
var ErrOrderMissingPaymentMethod = errors.New("order has no payment method")
