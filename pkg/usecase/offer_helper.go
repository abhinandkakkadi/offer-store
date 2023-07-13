package usecase

import (
	interfaces "github.com/abhinandkakkadi/offer-store/pkg/repository/interface"
	"github.com/abhinandkakkadi/offer-store/pkg/utils/models"
)


var OfferContainer = models.OfferCompany{}

func OfferUseCase(userRepository interfaces.UserRepository) {

	userRepository.GetOfferBasedOnCountry("US",OfferContainer)
	userRepository.GetOfferBasedOnCountry("CA",OfferContainer)

}
