package http

import (
	_ "github.com/abhinandkakkadi/offer-store/pkg/api/handler"
	handler "github.com/abhinandkakkadi/offer-store/pkg/api/handler"
	"github.com/abhinandkakkadi/offer-store/pkg/api/routes"
	interfaces "github.com/abhinandkakkadi/offer-store/pkg/repository/interface"
	"github.com/abhinandkakkadi/offer-store/pkg/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/swagger"
)

type ServerHTTP struct {
	engine *fiber.App
}

func NewServerHTTP(userHandler *handler.UserHandler) *ServerHTTP {

	router := fiber.New()

	router.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	router.Get("/swagger/*", swagger.HandlerDefault)

	routes.UserRoutes(router.Group("/"), userHandler)

	return &ServerHTTP{engine: router}
}

func (sh *ServerHTTP) Start(userRepository interfaces.UserRepository) {
	go usecase.OfferUseCase(userRepository)
	sh.engine.Listen(":3000")
}
