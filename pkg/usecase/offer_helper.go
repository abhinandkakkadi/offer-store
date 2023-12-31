package usecase

import (
	"time"

	interfaces "github.com/abhinandkakkadi/offer-store/pkg/repository/interface"
	"github.com/abhinandkakkadi/offer-store/pkg/utils/models"
)

// storage in main memory for efficient retrievel
var OfferContainers = map[string][]models.OfferCompany{}

// function to retrieve offers from database in every 10 seconds concurrently from every countries
func OfferUseCase(userRepository interfaces.UserRepository) {

	for {
		go func() {
			// OfferContainerUS = userRepository.GetOfferBasedOnCountry("US")
			OfferContainers["US"] = userRepository.GetOfferBasedOnCountry("US")
		}()

		go func() {
			// OfferContainerCA = userRepository.GetOfferBasedOnCountry("CA")
			OfferContainers["CA"] = userRepository.GetOfferBasedOnCountry("CA")
		}()

		go func() {
			// OfferContainerFR = userRepository.GetOfferBasedOnCountry("FR")
			OfferContainers["FR"] = userRepository.GetOfferBasedOnCountry("FR")
		}()

		go func() {
			// OfferContainerBR = userRepository.GetOfferBasedOnCountry("BR")
			OfferContainers["BR"] = userRepository.GetOfferBasedOnCountry("BR")
		}()

		time.Sleep(10 * time.Second)
	}

}

// checking the country for which offer should be returned
func ReturnOffer(country string) []models.OfferCompany {

	switch {
	case country == "US":
		return OfferContainers["US"]
	case country == "CA":
		return OfferContainers["CA"]
	case country == "FR":
		return OfferContainers["FR"]
	case country == "BR":
		return OfferContainers["BR"]
	}

	return []models.OfferCompany{}
}

var ValidCountryValues = []string{"US", "CA", "FR", "BR"}

// checking whether the country exists or not
func contains(value string) bool {

	for _, item := range ValidCountryValues {
		if item == value {
			return true
		}
	}
	return false

}
