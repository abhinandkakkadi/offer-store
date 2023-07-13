package repository

import (
	"fmt"

	interfaces "github.com/abhinandkakkadi/offer-store/pkg/repository/interface"
	"github.com/abhinandkakkadi/offer-store/pkg/utils/models"
	"gorm.io/gorm"
)

type UserDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &UserDatabase{DB}
}

func (u *UserDatabase) GetOfferBasedOnCountry(country string) []models.OfferCompany {

	// to check whether the offer still active and is after the the show_from date
	// var Offer []models.OfferCompany
	// if err := u.DB.Raw("SELECT * FROM offer_company WHERE country = ?	AND CURDATE() BETWEEN valid_from AND valid_to and CURDATE() >= show_from", country).Scan(&Offer).Error; err != nil {
	// 	fmt.Println(err)
	// }

	var Offer []models.OfferCompany
	if err := u.DB.Raw("SELECT * FROM offer_company WHERE country = ?	ORDER BY valid_from desc", country).Scan(&Offer).Error; err != nil {
		fmt.Println(err)
	}

	return Offer

}

func (u *UserDatabase) AddNewOffer(addOffer models.AddNewOffer) error {

	if err := u.DB.Exec(`
	INSERT INTO offer_company (client_id, country, image, image_width, image_height, text_locale, validity_text_locale, position, valid_from, show_from, valid_to, flag, page_count, store_url, store_url_title, offer_home) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)
`, addOffer.ClientID, addOffer.Country, addOffer.Image, addOffer.ImageWidth, addOffer.ImageHeight, addOffer.TextLocale, addOffer.ValidityTextLocale, addOffer.Position, addOffer.ValidFrom, addOffer.ShowFrom, addOffer.ValidTo, addOffer.Flag, addOffer.PageCount, addOffer.StoreURL, addOffer.StoreURLTitle, addOffer.OfferHome).Error; err != nil {
		return err
	}

	return nil

}
