package order

type OrderService struct {
	pApi OrderApi
}

func NewOrderService(pApi OrderApi) *OrderService {
	return &OrderService{
		pApi: pApi,
	}
}

func (s *OrderService) GetManyOrders(in *GetManyOrdersIn) (*GetManyOrdersOut, error) {
	return s.pApi.GetManyOrders(in)
}

func (s *OrderService) GetOrder(id string) (*Order, error) {
	return s.pApi.GetOrder(id)
}

func (s *OrderService) CreateOrder(in *CreateOrderIn) (*Order, error) {
	return s.pApi.CreateOrder(in)
}
