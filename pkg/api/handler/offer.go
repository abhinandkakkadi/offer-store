package handler

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	services "github.com/abhinandkakkadi/offer-store/pkg/usecase/interface"
	"github.com/abhinandkakkadi/offer-store/pkg/utils/models"
	"github.com/abhinandkakkadi/offer-store/pkg/utils/response"
)

type UserHandler struct {
	userUseCase services.UserUseCase
}

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

// @Summary Get Offer Based On Country
// @Description Get Latest Offer from different Countries by passing country short form
// @Tags OFFERS
// @Accept json
// @Produce json
// @Param country path string true "country short form"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /get-offer/{country} [get]
func (u *UserHandler) GetValueBasedOnCountry(c *fiber.Ctx) error {

	country := c.Params("country")
	offerBasedOnCountry, err := u.userUseCase.GetValueBasedOnCountry(country)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "field not in right format", nil, err.Error())
		c.Status(http.StatusBadRequest).JSON(errRes)
		return nil
	}

	sucRes := response.ClientResponse(http.StatusOK, "successfully retrieved offers", offerBasedOnCountry, nil)

	contentType := c.Get("Type")

	if contentType == "json" {

		c.JSON(sucRes)

	} else if contentType == "bson" {

		bsonData, err := bson.Marshal(sucRes)
		if err != nil {
			c.Status(fiber.StatusInternalServerError).SendString("Error marshaling data to BSON")
			return nil
		}
		c.Set(fiber.HeaderContentType, "application/bson")
		c.Send(bsonData)

	} else {

		c.Status(http.StatusBadRequest).SendString("specify json or bson")

	}

	// c.Status(http.StatusOK).JSON(sucRes)

	return nil
}

// @Summary Add A new Offer for an existing Country
// @Description Add a new offer by passing in the whole offer details including offer validity date
// @Tags OFFERS
// @Accept json
// @Produce json
// @Param user body models.AddNewOffer true "Add a new offer"
// @Success 200 {object} response.Response{}
// @Failure 500 {object} response.Response{}
// @Router /add-offer [post]
func (u *UserHandler) AddNewOffer(c *fiber.Ctx) error {

	var addOffer models.AddNewOffer
	if err := c.BodyParser(&addOffer); err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "field not in right format", nil, err.Error())
		c.Status(http.StatusBadRequest).JSON(errRes)
		return nil
	}

	fmt.Println(addOffer)

	err := u.userUseCase.AddNewOffer(addOffer)
	if err != nil {
		errRes := response.ClientResponse(http.StatusBadRequest, "could not add offer", nil, err.Error())
		c.Status(http.StatusBadRequest).JSON(errRes)
		return nil
	}

	sucRes := response.ClientResponse(http.StatusCreated, "successfully added new offer", nil, nil)
	c.Status(http.StatusCreated).JSON(sucRes)

	return nil

}
