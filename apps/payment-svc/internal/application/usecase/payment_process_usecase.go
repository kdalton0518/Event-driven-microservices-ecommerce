package usecase

import (
	"github.com/buemura/event-driven-commerce/payment-svc/internal/domain/payment"
	"github.com/buemura/event-driven-commerce/payment-svc/internal/infra/util"
)

type PaymentProcessUsecase struct {
	repo payment.PaymentRepository
}

func NewPaymentProcessUsecase(repo payment.PaymentRepository) *PaymentProcessUsecase {
	return &PaymentProcessUsecase{
		repo: repo,
	}
}

func (u *PaymentProcessUsecase) Execute(in *payment.ProcessPaymentIn) (*payment.Payment, error) {
	p, err := u.repo.FindPendingByOrderId(in.OrderId)
	if err != nil {
		return nil, err
	}
	if p == nil {
		return nil, payment.ErrNoPendingPaymentFound
	}

	p.Status = u.setPaymentStatus()

	err = u.repo.Update(p.ID, string(p.Status))
	if err != nil {
		return nil, err
	}

	return p, nil
}

func (u *PaymentProcessUsecase) setPaymentStatus() payment.PaymentStatus {
	pick := util.RandomNumber()

	var status payment.PaymentStatus

	switch pick {
	case 1:
		status = payment.PaymentPaid
	case 2:
		status = payment.PaymentFailed
	}

	return status
}
