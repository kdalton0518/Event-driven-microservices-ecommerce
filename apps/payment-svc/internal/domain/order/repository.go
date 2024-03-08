package order

type OrderRepositoryPaginatedOut struct {
	OrderList  []*Order
	TotalCount int
}

type OrderRepository interface {
	FindById(string) (*Order, error)
	Save(*Order) (*Order, error)
}
