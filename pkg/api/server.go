package http

import (
	_ "github.com/abhinandkakkadi/offer-store/pkg/api/handler"
	handler "github.com/abhinandkakkadi/offer-store/pkg/api/handler"
	"github.com/abhinandkakkadi/offer-store/pkg/api/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler) *ServerHTTP {
	router := gin.New()

	router.LoadHTMLGlob("templates/*.html")

	router.Use(gin.Logger())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.UserRoutes(router.Group("/"), userHandler)

	return &ServerHTTP{engine: router}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
