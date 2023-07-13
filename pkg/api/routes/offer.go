package routes

import (
	"github.com/abhinandkakkadi/offer-store/pkg/api/handler"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router, userHandler *handler.UserHandler) {

	router.Post("/signup", userHandler.UserSignUp)
	router.Post("/login", userHandler.LoginHandler)

}
