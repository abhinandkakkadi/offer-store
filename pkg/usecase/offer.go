package usecase

import interfaces "github.com/abhinandkakkadi/offer-store/pkg/repository/interface"

func OfferUseCase(userRepository interfaces.UserRepository) {

	userRepository.GetOfferBasedOnCountry()

}
