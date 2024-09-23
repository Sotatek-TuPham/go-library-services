package routes

import (
	"library-server/handler"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.RouterGroup) {
	r.POST("/login", handler.Login)
}
