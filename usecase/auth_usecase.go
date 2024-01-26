package usecase

import (
	"go-merchant/entity/dto"
	"go-merchant/shared/service"
)

type AuthUsecase interface {
	Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error)
}

type authUsecase struct {
	customerUC CustomerUsecase
	jwtService service.JwtService
}

// Login implements AuthUsecase.
func (a *authUsecase) Login(payload dto.AuthRequestDto) (dto.AuthResponseDto, error) {
	customer, err := a.customerUC.FindCustomerForLogin(payload.Username, payload.Password)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}
	token, err := a.jwtService.CreateToken(customer)
	if err != nil {
		return dto.AuthResponseDto{}, err
	}

	return token, nil
}

func NewAuthUsecase(customerUC CustomerUsecase, jwtService service.JwtService) AuthUsecase {
	return &authUsecase{customerUC: customerUC, jwtService: jwtService}
}
