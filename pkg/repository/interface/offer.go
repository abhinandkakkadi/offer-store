package interfaces

import "github.com/abhinandkakkadi/offer-store/pkg/utils/models"

type UserRepository interface {
	GetOfferBasedOnCountry(country string) []models.OfferCompany
	AddNewOffer(addOffer models.AddNewOffer) error
}
