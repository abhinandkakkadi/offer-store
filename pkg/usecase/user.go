package usecase

import (
	services "github.com/abhinandkakkadi/offer-store/pkg/usecase/interface"
)

type userUseCase struct {
}

func NewUserUseCase() services.UserUseCase {
	return &userUseCase{}
}

func (u *userUseCase) UserSignUp() {

}

func (u *userUseCase) LoginHandler() {

}
