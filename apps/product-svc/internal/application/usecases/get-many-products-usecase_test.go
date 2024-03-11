package usecases

import (
	"testing"

	"github.com/buemura/event-driven-commerce/product-svc/internal/domain/common"
	"github.com/buemura/event-driven-commerce/product-svc/internal/domain/product"
	"github.com/buemura/event-driven-commerce/product-svc/internal/infra/database"
	"github.com/stretchr/testify/assert"
)

func TestGetManyProductsUsecase(t *testing.T) {
	p := []*product.Product{
		{
			ID:          1,
			Name:        "existing_product",
			Description: "existing_product_description",
			Price:       100,
			Quantity:    14,
			ImageUrl:    "http://localhost:3131",
		},
		{
			ID:          2,
			Name:        "existing_product_2",
			Description: "existing_product_description_2",
			Price:       433,
			Quantity:    122,
			ImageUrl:    "http://localhost:3131",
		},
		{
			ID:          3,
			Name:        "existing_product_3",
			Description: "existing_product_description_3",
			Price:       113,
			Quantity:    3232,
			ImageUrl:    "http://localhost:3131",
		},
	}

	repo := database.NewInMemoryProductRepo(p)
	service := NewGetManyProductUsecase(repo)

	t.Run("Return product slice", func(t *testing.T) {
		res, _ := service.Execute(&product.GetManyProductsIn{
			Page:  1,
			Items: 10,
		})
		assert.Equal(t, p, res.ProductList)
		assert.Equal(t, &common.PaginationMeta{
			Page:       1,
			Items:      10,
			TotalPages: 1,
			TotalItems: len(res.ProductList),
		}, res.Meta)
	})
}
