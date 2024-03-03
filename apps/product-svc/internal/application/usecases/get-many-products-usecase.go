package usecases

import (
	"math"

	"github.com/buemura/event-driven-commerce/product-svc/internal/domain/common"
	"github.com/buemura/event-driven-commerce/product-svc/internal/domain/product"
)

type GetManyProductUsecase struct {
	repo product.ProductRepository
}

func NewGetManyProductUsecase(repo product.ProductRepository) *GetManyProductUsecase {
	return &GetManyProductUsecase{
		repo: repo,
	}
}

func (uc *GetManyProductUsecase) Execute(opt *product.GetManyProductsIn) (*product.GetManyProductsOut, error) {
	res, err := uc.repo.FindMany()
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(res.TotalCount) / float64(opt.Items)))

	return &product.GetManyProductsOut{
		ProductList: res.ProductList,
		Meta: &common.PaginationMeta{
			Page:       opt.Page,
			Items:      opt.Items,
			TotalPages: totalPages,
			TotalItems: res.TotalCount,
		},
	}, nil
}
