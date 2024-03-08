package payment

type PaymentStatus string

const (
	PaymentPending PaymentStatus = "PENDING"
	PaymentPaid    PaymentStatus = "PAID"
	PaymentFailed  PaymentStatus = "FAILED"
)

type PaymentMethod string

const (
	CreditCard PaymentMethod = "CREDIT_CARD"
	Pix        PaymentMethod = "PIX"
)
