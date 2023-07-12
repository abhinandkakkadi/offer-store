package handler

import (
	"github.com/gofiber/fiber"

	services "github.com/abhinandkakkadi/offer-store/pkg/usecase/interface"
)

type UserHandler struct {
	userUseCase services.UserUseCase
}

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

// @Summary SignUp functionality for user
// @Description SignUp functionality at the user side
// @Tags User Authentication
// @Accept json
// @Produce json
// @Param user body models.UserDetails true "User Details Input"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /signup [post]
func (u *UserHandler) UserSignUp(c *fiber.Ctx) {



}

// @Summary LogIn functionality for user
// @Description LogIn functionality at the user side
// @Tags User Authentication
// @Accept json
// @Produce json
// @Param user body models.UserLogin true "User Details Input"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /login [post]
func (u *UserHandler) LoginHandler(c *fiber.Ctx) {


}
