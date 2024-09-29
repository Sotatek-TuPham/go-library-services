package routes

import (
	"order-server/handler"
	"order-server/service"
	"os"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	baseURL := os.Getenv("LIBRARY_SERVER_URL")
	userService := service.NewUserService(baseURL)
	userHandler := handler.NewUserHandler(userService)

	router.POST("/receipts", userHandler.PlaceReceipt)
	router.POST("/receipts/:id/cancel", userHandler.CancelReceipt)
	router.GET("/receipts", userHandler.GetReceiptsByUserID)
}
