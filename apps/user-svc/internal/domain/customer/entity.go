package customer

import (
	"errors"

	"github.com/google/uuid"
)

type Customer struct {
	ID       string
	Name     string
	Email    string
	Password string
}

func NewCustomer(input *CreateCustomerIn) (*Customer, error) {
	err := validate(input)
	if err != nil {
		return nil, err
	}

	return &Customer{
		ID:       uuid.NewString(),
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}, nil
}

func validate(input *CreateCustomerIn) error {
	if input.Name == "" {
		return errors.New("invalid customer name")
	}

	if input.Email == "" {
		return errors.New("invalid customer email")
	}

	if len(input.Password) < 6 {
		return errors.New("invalid customer password")
	}

	return nil
}
