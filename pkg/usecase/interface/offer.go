package interfaces

import "github.com/abhinandkakkadi/offer-store/pkg/utils/models"

type UserUseCase interface {
	GetValueBasedOnCountry(country string) ([]models.OfferCompany,error)
	AddNewOffer(addOffer models.AddNewOffer) error
}
