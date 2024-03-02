package usecases

import "github.com/buemura/event-driven-commerce/customer-svc/internal/domain/customer"

type CustomerGetService struct {
	repo customer.CustomerRepository
}

func NewCustomerGetService(repo customer.CustomerRepository) *CustomerGetService {
	return &CustomerGetService{
		repo: repo,
	}
}

func (s *CustomerGetService) Execute(customerID string) (*customer.CustomerOut, error) {
	cust, err := s.repo.FindById(customerID)
	if err != nil {
		return nil, err
	}
	if cust == nil {
		return nil, customer.ErrCustomerNotFound
	}

	return &customer.CustomerOut{
		ID:    cust.ID,
		Name:  cust.Name,
		Email: cust.Email,
	}, nil
}
