package payment

import "errors"

var (
	ErrCannotCreateWithoutOrderId = errors.New("cannot create payment without orderId")
	ErrNoPendingPaymentFound      = errors.New("no pending payment order found")
)
