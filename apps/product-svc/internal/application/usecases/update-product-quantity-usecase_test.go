package usecases

import (
	"testing"

	"github.com/buemura/event-driven-commerce/product-svc/internal/domain/product"
	"github.com/buemura/event-driven-commerce/product-svc/internal/infra/database"
	"github.com/stretchr/testify/assert"
)

func TestUpdateProductQuantityUsecase(t *testing.T) {
	p := []*product.Product{
		{
			ID:       1,
			Name:     "existing_product",
			Price:    100,
			Quantity: 14,
			ImageUrl: "http://localhost:3131",
		},
		{
			ID:       2,
			Name:     "existing_product_2",
			Price:    433,
			Quantity: 122,
			ImageUrl: "http://localhost:3131",
		},
		{
			ID:       3,
			Name:     "existing_product_3",
			Price:    113,
			Quantity: 3232,
			ImageUrl: "http://localhost:3131",
		},
	}

	repo := database.NewInMemoryProductRepo(p)
	service := NewUpdateProductQuantityUsecase(repo)

	t.Run("Return product struct", func(t *testing.T) {
		res, _ := service.Execute(&product.UpdateProductQuantityIn{
			ID:       3,
			Quantity: -2,
		})
		assert.Equal(t, &product.Product{
			ID:       3,
			Name:     "existing_product_3",
			Price:    113,
			Quantity: 3230,
			ImageUrl: "http://localhost:3131",
		}, res)
	})

	t.Run("Return ErrProductInsufficientQuantity when product has no enough quantity", func(t *testing.T) {
		_, err := service.Execute(&product.UpdateProductQuantityIn{
			ID:       1,
			Quantity: -15,
		})

		assert.Equal(t, err, product.ErrProductInsufficientQuantity)
		assert.NotNil(t, err)
	})

	t.Run("Return ErrProductNotFound when product does not exists", func(t *testing.T) {
		_, err := service.Execute(&product.UpdateProductQuantityIn{
			ID:       99,
			Quantity: -15,
		})

		assert.Equal(t, err, product.ErrProductNotFound)
		assert.NotNil(t, err)
	})
}
