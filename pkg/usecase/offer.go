package usecase

import (
	"errors"

	interfaces "github.com/abhinandkakkadi/offer-store/pkg/repository/interface"
	services "github.com/abhinandkakkadi/offer-store/pkg/usecase/interface"
	"github.com/abhinandkakkadi/offer-store/pkg/utils/models"
)

type userUseCase struct {
	repo interfaces.UserRepository
}

func NewUserUseCase(repo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		repo: repo,
	}
}

func (u *userUseCase) GetValueBasedOnCountry(country string) ([]models.OfferCompany, error) {

	countryExist := contains(country)
	if !countryExist {
		return []models.OfferCompany{}, errors.New("country does not exists")
	}

	offer := ReturnOffer(country)
	return offer, nil

}

func (u *userUseCase) AddNewOffer(addOffer models.AddNewOffer) error {

	countryExist := contains(addOffer.Country)
	if !countryExist {
		return errors.New("country does not exists")
	}

	err := u.repo.AddNewOffer(addOffer)
	if err != nil {
		return err
	}

	return nil

}
