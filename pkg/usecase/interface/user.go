package interfaces

import (
	"github.com/abhinandkakkadi/offer-store/pkg/utils/models"
)

type UserUseCase interface {
	UserSignUp(user models.UserDetails) (models.TokenUsers, error)
	LoginHandler(user models.UserLogin) (models.TokenUsers, error)
}
