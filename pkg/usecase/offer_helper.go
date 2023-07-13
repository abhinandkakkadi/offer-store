package usecase

import (
	interfaces "github.com/abhinandkakkadi/offer-store/pkg/repository/interface"
	"github.com/abhinandkakkadi/offer-store/pkg/utils/models"
)


var OfferContainerUS = models.OfferCompany{}
var OfferContainerCA = models.OfferCompany{}

func OfferUseCase(userRepository interfaces.UserRepository) {

	OfferContainerUS = userRepository.GetOfferBasedOnCountry("US",OfferContainerUS)
	OfferContainerCA = userRepository.GetOfferBasedOnCountry("CA",OfferContainerCA)
	u := userUseCase{}
	u.GetValueBasedOnCountry()

}

func ReturnOffer(country string) models.OfferCompany {

	if country == "US" {
		return OfferContainerUS
	}

	if country == "CA" {
		return OfferContainerCA
	}

	return models.OfferCompany{}
}
