package routes

import (
	"library-server/handler"

	"github.com/gin-gonic/gin"
)

func ReceiptRoutes(router *gin.RouterGroup) {
	router.POST("/", handler.CreateReceipt)
	router.GET("/:id", handler.GetReceiptByID)
	router.GET("/user/:user_id", handler.GetReceiptsByUserID)
	router.PATCH("/:id/status", handler.UpdateReceiptStatus)
	router.DELETE("/:id", handler.DeleteReceipt)
	router.GET("/", handler.GetAllReceipts)
}
