package usecase

import (
	"go-merchant/entity"
	"go-merchant/repository"
)

type CustomerUsecase interface {
	FindCustomerForLogin(username, password string) (entity.Customer, error)
}

// hubungkan ke repository
type customerUsecase struct {
	repo repository.CustomerRepository
}

// FindCustomerForLogin implements CustomerUsecase.
func (c *customerUsecase) FindCustomerForLogin(username, password string) (entity.Customer, error) {
	return c.repo.GetByUsernamePassword(username, password)
}

// buat bridge
func NewCustomerUsecase(repo repository.CustomerRepository) CustomerUsecase {
	return &customerUsecase{repo: repo}
}
