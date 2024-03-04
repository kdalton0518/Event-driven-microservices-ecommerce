package order

type OrderRepositoryPaginatedOut struct {
	OrderList  []*Order
	TotalCount int
}

type OrderRepository interface {
	FindMany(*GetManyOrdersIn) (*OrderRepositoryPaginatedOut, error)
	FindById(string) (*Order, error)
	Save(*Order) (*Order, error)
}
