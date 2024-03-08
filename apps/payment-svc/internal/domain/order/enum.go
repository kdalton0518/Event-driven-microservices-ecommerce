package order

type OrderStatus string

const (
	StatusPending   OrderStatus = "PENDING"
	StatusCompleted OrderStatus = "COMPLETED"
	StatusCancelled OrderStatus = "CANCELLED"
)
