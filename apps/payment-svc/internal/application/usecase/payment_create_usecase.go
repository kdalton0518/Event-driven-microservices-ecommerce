package usecase

import (
	"github.com/buemura/event-driven-commerce/payment-svc/internal/domain/payment"
)

type PaymentCreateUsecase struct {
	repo payment.PaymentRepository
}

func NewPaymentCreateUsecase(repo payment.PaymentRepository) *PaymentCreateUsecase {
	return &PaymentCreateUsecase{
		repo: repo,
	}
}

func (u *PaymentCreateUsecase) Execute(in *payment.CreatePaymentIn) (*payment.Payment, error) {
	p, err := payment.NewPayment(in)
	if err != nil {
		return nil, err
	}

	_, err = u.repo.Save(p)
	if err != nil {
		return nil, err
	}

	return p, err
}
