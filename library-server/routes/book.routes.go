package routes

import (
	"library-server/handler"
	"library-server/middleware"

	"github.com/gin-gonic/gin"
)

func BookRoutes(router *gin.RouterGroup) {
	router.Use(middleware.Authenticate())
	router.POST("/", handler.CreateBook)
	router.GET("/:id", handler.GetBookByID)
	router.GET("/", handler.GetAllBooks)
	router.PUT("/:id", handler.UpdateBook)
	router.DELETE("/:id", handler.DeleteBook)
	router.GET("/category/:categoryID", handler.GetBooksByCategory)
	router.PATCH("/:id/status", handler.UpdateBookStatus)
}
