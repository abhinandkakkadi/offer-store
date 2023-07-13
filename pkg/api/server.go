package http

import (
	_ "github.com/abhinandkakkadi/offer-store/pkg/api/handler"
	handler "github.com/abhinandkakkadi/offer-store/pkg/api/handler"
	"github.com/abhinandkakkadi/offer-store/pkg/api/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

type ServerHTTP struct {
	engine *fiber.App
}

func NewServerHTTP(userHandler *handler.UserHandler) *ServerHTTP {

	router := fiber.New()

	router.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	routes.UserRoutes(router.Group("/"), userHandler)

	return &ServerHTTP{engine: router}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Listen(":3000")
}
