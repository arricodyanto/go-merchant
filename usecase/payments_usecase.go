package usecase

import (
	"fmt"
	"go-merchant/entity"
	"go-merchant/repository"
)

type PaymentUsecase interface {
	CreatePayment(payload entity.Payment) (entity.Payment, error)
	UpdateStatusPayment(id string) (entity.Payment, error)
}

type paymentUsecase struct {
	repo repository.PaymentRepository
}

// CreatePayment implements PaymentUsecase.
func (p *paymentUsecase) CreatePayment(payload entity.Payment) (entity.Payment, error) {
	// cek required fields
	if payload.CustomerID == "" || payload.Amount == 0 {
		return entity.Payment{}, fmt.Errorf("oops, filed required")
	}

	// eksekusi repo
	payment, err := p.repo.Create(payload)
	if err != nil {
		return entity.Payment{}, fmt.Errorf("failed to create new payment: %v", err.Error())
	}

	// return value jika berhasil
	return payment, nil
}

// UpdateStatusPayment implements PaymentUsecase.
func (p *paymentUsecase) UpdateStatusPayment(id string) (entity.Payment, error) {
	panic("unimplemented")
}

func NewPaymentUsecase(repo repository.PaymentRepository) PaymentUsecase {
	return &paymentUsecase{repo: repo}
}
