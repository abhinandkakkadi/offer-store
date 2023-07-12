package routes

import (
	"github.com/abhinandkakkadi/offer-store/pkg/api/handler"
	"github.com/abhinandkakkadi/offer-store/pkg/api/middleware"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router, userHandler *handler.UserHandler) {

	router.Use(middleware.AuthMiddleware())
	router.Post("/signup", userHandler.UserSignUp)
	router.Post("/login", userHandler.LoginHandler)
	
}
