package routes

import (
	"github.com/abhinandkakkadi/offer-store/pkg/api/handler"
	"github.com/abhinandkakkadi/offer-store/pkg/api/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup, userHandler *handler.UserHandler) {

	router.Use(middleware.AuthMiddleware())
	router.POST("/signup", userHandler.UserSignUp)
	router.POST("/login", userHandler.LoginHandler)
	
}
