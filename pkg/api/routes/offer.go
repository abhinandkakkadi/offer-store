package routes

import (
	"github.com/abhinandkakkadi/offer-store/pkg/api/handler"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(router fiber.Router, userHandler *handler.UserHandler) {

	router.Get("/get-offer/:country", userHandler.GetValueBasedOnCountry)
	router.Post("/add-offer", userHandler.AddNewOffer)

}
