package application

import (
	"user-svc/internal/domain/customer"
	"user-svc/internal/domain/encryption"
)

type CustomerSigninService struct {
	repo           customer.CustomerRepository
	passwordHasher encryption.PasswordHasher
}

func NewCustomerSigninService(
	repo customer.CustomerRepository,
	passwordHasher encryption.PasswordHasher,
) *CustomerSigninService {
	return &CustomerSigninService{
		repo:           repo,
		passwordHasher: passwordHasher,
	}
}

func (s *CustomerSigninService) Execute(in *customer.SigninCustomerIn) (*customer.SigninCustomerOut, error) {
	cust, err := s.repo.FindByEmail(in.Email)
	if err != nil {
		return nil, err
	}
	if cust == nil {
		return nil, customer.ErrCustomerInvalidCredential
	}

	isPassValid, err := s.passwordHasher.Comapre(in.Password, cust.Password)
	if err != nil {
		return nil, err
	}
	if !isPassValid {
		return nil, customer.ErrCustomerInvalidCredential
	}

	return &customer.SigninCustomerOut{
		AccessToken: "token",
		Customer: customer.CustomerOut{
			ID:    cust.ID,
			Name:  cust.Name,
			Email: cust.Email,
		},
	}, nil
}
