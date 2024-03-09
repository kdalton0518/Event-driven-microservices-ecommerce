package payment

type PaymentRepository interface {
	FindById(string) (*Payment, error)
	FindByOrderId(string) ([]*Payment, error)
	FindPendingByOrderId(string) (*Payment, error)
	Save(*Payment) (*Payment, error)
	Update(id, status string) error
}
