package http

import (
	_ "github.com/abhinandkakkadi/offer-store/pkg/api/handler"
	handler "github.com/abhinandkakkadi/offer-store/pkg/api/handler"
	"github.com/abhinandkakkadi/offer-store/pkg/api/routes"
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber"
	"gorm.io/gorm/logger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler) *ServerHTTP {
	
	router := fiber.New()
	router.Use(logger.New(logger.Writer, logger.Config))
	
	routes.UserRoutes(router.Group("/"), userHandler)

	return &ServerHTTP{engine: router}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
