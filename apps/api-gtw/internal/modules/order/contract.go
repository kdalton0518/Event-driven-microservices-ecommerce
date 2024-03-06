package order

type OrderApi interface {
	GetManyOrders(*GetManyOrdersIn) (*GetManyOrdersOut, error)
	GetOrder(string) (*Order, error)
	CreateOrder(*CreateOrderIn) (*Order, error)
}
