package repository

import (
	interfaces "github.com/abhinandkakkadi/offer-store/pkg/repository/interface"
	"gorm.io/gorm"
)

type UserDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &UserDatabase{DB}
}

func (u *UserDatabase) GetOfferBasedOnCountry() {

}
