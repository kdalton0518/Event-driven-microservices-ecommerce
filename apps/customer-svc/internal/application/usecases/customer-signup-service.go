package usecases

import (
	"github.com/buemura/event-driven-commerce/customer-svc/internal/application/contracts"
	"github.com/buemura/event-driven-commerce/customer-svc/internal/domain/customer"
)

type CustomerSignupService struct {
	repo           customer.CustomerRepository
	passwordHasher contracts.PasswordHasher
}

func NewCustomerSignupService(
	repo customer.CustomerRepository,
	passwordHasher contracts.PasswordHasher,
) *CustomerSignupService {
	return &CustomerSignupService{
		repo:           repo,
		passwordHasher: passwordHasher,
	}
}

func (s *CustomerSignupService) Execute(in *customer.CreateCustomerIn) error {
	custExists, err := s.repo.FindByEmail(in.Email)
	if err != nil {
		return err
	}
	if custExists != nil {
		return customer.ErrCustomerAlreadyExists
	}

	hashed, err := s.passwordHasher.Hash(in.Password)
	if err != nil {
		return err
	}

	in.Password = hashed

	cust, err := customer.NewCustomer(in)
	if err != nil {
		return err
	}

	_, err = s.repo.Save(cust)
	if err != nil {
		return err
	}
	return nil
}
