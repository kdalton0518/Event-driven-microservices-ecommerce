package usecase

import (
	"log"

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
	log.Println("[PaymentCreateUsecase][Execute] - Init payment creation for order:", in.OrderId)
	p, err := payment.NewPayment(in)
	if err != nil {
		return nil, err
	}

	_, err = u.repo.Save(p)
	if err != nil {
		return nil, err
	}

	log.Println("[PaymentCreateUsecase][Execute] - Successfully inserted payment for order:", in.OrderId)
	return p, err
}
