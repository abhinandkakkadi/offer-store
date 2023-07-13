package usecase

import (
	"time"

	interfaces "github.com/abhinandkakkadi/offer-store/pkg/repository/interface"
	"github.com/abhinandkakkadi/offer-store/pkg/utils/models"
)

// RAM storage for all countries
var OfferContainerUS = []models.OfferCompany{}
var OfferContainerCA = []models.OfferCompany{}
var OfferContainerFR = []models.OfferCompany{}
var OfferContainerBR = []models.OfferCompany{}


// function to retrieve offers from database in every 10 seconds 
func OfferUseCase(userRepository interfaces.UserRepository) {

	for {
		go func() {
			OfferContainerUS = userRepository.GetOfferBasedOnCountry("US")
		}()

		go func() {
			OfferContainerCA = userRepository.GetOfferBasedOnCountry("CA")
		}()

		go func() {
			OfferContainerFR = userRepository.GetOfferBasedOnCountry("FR")
		}()

		go func() {
			OfferContainerBR = userRepository.GetOfferBasedOnCountry("BR")
		}()

		time.Sleep(10 * time.Second)
	}

}

// checking the country for which offer should be returned
func ReturnOffer(country string) []models.OfferCompany {

	switch {
	case country == "US":
		return OfferContainerUS
	case country == "CA":
		return OfferContainerCA
	case country == "FR":
		return OfferContainerFR
	case country == "BR":
		return OfferContainerBR
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
