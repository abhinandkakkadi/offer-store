package usecase

import (
	"fmt"

	services "github.com/abhinandkakkadi/offer-store/pkg/usecase/interface"
)

type userUseCase struct {
}

func NewUserUseCase() services.UserUseCase {
	return &userUseCase{}
}

func (u *userUseCase) GetValueBasedOnCountry() {

	offer := ReturnOffer("US")
	fmt.Println(offer)
}

func (u *userUseCase) LoginHandler() {

}
