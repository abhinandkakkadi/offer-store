package http

import (
	"compress/gzip"

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

		// Middleware function for data compression
		router.Use(compress.New(compress.Config{
			Level: gzip.BestSpeed,
	}))

	// Setting Swagger Routes
	router.Get("/swagger/*", swagger.HandlerDefault)

	routes.UserRoutes(router.Group("/"), userHandler)

	return &ServerHTTP{engine: router}
}

// server starting
func (sh *ServerHTTP) Start(userRepository interfaces.UserRepository) {
	go usecase.OfferUseCase(userRepository)
	sh.engine.Listen(":3000")
}
