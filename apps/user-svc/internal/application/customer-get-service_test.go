package application

import (
	"testing"
	"user-svc/internal/domain/customer"
	"user-svc/internal/infra/database"

	"github.com/stretchr/testify/assert"
)

func TestCustomerGetService(t *testing.T) {

	c := []customer.Customer{
		{
			ID:       "existing_id",
			Name:     "existing_name",
			Email:    "existing_email",
			Password: "hashed:existing_password",
		},
	}

	repo := database.NewInMemoryCustomerRepo(c)
	service := NewCustomerGetService(repo)

	t.Run("Customer not found", func(t *testing.T) {
		_, err := service.Execute("not_found_id")
		assert.Equal(t, err, customer.ErrCustomerNotFound)
		assert.NotNil(t, err)
	})

	t.Run("Customer get success", func(t *testing.T) {
		res, _ := service.Execute("existing_id")
		assert.Equal(t, &customer.CustomerOut{
			ID:    "existing_id",
			Name:  "existing_name",
			Email: "existing_email",
		}, res)
	})
}
