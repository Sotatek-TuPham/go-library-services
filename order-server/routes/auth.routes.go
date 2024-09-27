package routes

import (
	"order-server/handler"
	"order-server/service"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	authHandler := handler.NewAuthHandler(service.NewAuthService())

	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", authHandler.Register)
		authGroup.POST("/login", authHandler.Login)
	}
}
