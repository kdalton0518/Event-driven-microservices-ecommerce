package usecases

import (
	"customer-svc/internal/application/contracts"
	"customer-svc/internal/domain/customer"
)

type CustomerSigninService struct {
	repo           customer.CustomerRepository
	passwordHasher contracts.PasswordHasher
	tokenGenerator contracts.TokenGenerator
}

func NewCustomerSigninService(
	repo customer.CustomerRepository,
	passwordHasher contracts.PasswordHasher,
	tokenGenerator contracts.TokenGenerator,
) *CustomerSigninService {
	return &CustomerSigninService{
		repo:           repo,
		passwordHasher: passwordHasher,
		tokenGenerator: tokenGenerator,
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

	isPassValid := s.passwordHasher.Compare(in.Password, cust.Password)
	if !isPassValid {
		return nil, customer.ErrCustomerInvalidCredential
	}

	accessToken, err := s.tokenGenerator.Generate(cust.ID)
	if err != nil {
		return nil, err
	}

	return &customer.SigninCustomerOut{
		AccessToken: accessToken,
		Customer: customer.CustomerOut{
			ID:    cust.ID,
			Name:  cust.Name,
			Email: cust.Email,
		},
	}, nil
}
