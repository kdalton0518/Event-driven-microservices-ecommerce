package usecases

import (
	"math"

	"github.com/buemura/event-driven-commerce/order-svc/internal/domain/common"
	"github.com/buemura/event-driven-commerce/order-svc/internal/domain/order"
)

type GetManyOrdersUsecase struct {
	repo order.OrderRepository
}

func NewGetManyOrdersUsecase(repo order.OrderRepository) *GetManyOrdersUsecase {
	return &GetManyOrdersUsecase{
		repo: repo,
	}
}

func (uc *GetManyOrdersUsecase) Execute(opt *order.GetManyOrdersIn) (*order.GetManyOrdersOut, error) {
	res, err := uc.repo.FindMany(opt)
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(res.TotalCount) / float64(opt.Items)))

	return &order.GetManyOrdersOut{
		OrderList: res.OrderList,
		Meta: &common.PaginationMeta{
			Page:       opt.Page,
			Items:      opt.Items,
			TotalPages: totalPages,
			TotalItems: res.TotalCount,
		},
	}, nil
}
