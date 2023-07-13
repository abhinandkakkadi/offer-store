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

func (u *UserDatabase) GetOfferBasedOnCountry(country string,OfferContainer models.OfferCompany) models.OfferCompany {

	if err := u.DB.Raw("select * from offer_company where country = ?",country).Scan(&OfferContainer).Error; err != nil  {
		fmt.Println(err)
	}	

	return OfferContainer
	
}
