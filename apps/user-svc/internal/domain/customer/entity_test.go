package customer

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomer(t *testing.T) {
	t.Run("invalid name", func(t *testing.T) {
		input := &CreateCustomerIn{
			Name:     "",
			Email:    "any_email",
			Password: "any_password",
		}

		_, err := NewCustomer(input)
		assert.Equal(t, err, errors.New("invalid customer name"))
		assert.NotNil(t, err)
	})

	t.Run("invalid email", func(t *testing.T) {
		input := &CreateCustomerIn{
			Name:     "any_name",
			Email:    "",
			Password: "any_password",
		}

		_, err := NewCustomer(input)
		assert.Equal(t, err, errors.New("invalid customer email"))
		assert.NotNil(t, err)
	})

	t.Run("invalid password", func(t *testing.T) {
		input := &CreateCustomerIn{
			Name:     "any_name",
			Email:    "any_email",
			Password: "",
		}

		_, err := NewCustomer(input)
		assert.Equal(t, err, errors.New("invalid customer password"))
		assert.NotNil(t, err)
	})
}
