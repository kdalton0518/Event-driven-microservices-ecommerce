package usecases

import (
	"testing"

	"github.com/buemura/event-driven-commerce/product-svc/internal/domain/product"
	"github.com/buemura/event-driven-commerce/product-svc/internal/infra/database"
	"github.com/stretchr/testify/assert"
)

func TestGetManyProductsUsecase(t *testing.T) {
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
	service := NewGetManyProductUsecase(repo)

	t.Run("Return product slice", func(t *testing.T) {
		res, _ := service.Execute()
		assert.Equal(t, p, res)
	})
}
