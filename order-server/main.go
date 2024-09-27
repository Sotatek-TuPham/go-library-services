package main

import (
	"os"

	db "order-server/DB"
	_ "order-server/docs" // Import swagger docs
	"order-server/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Order API
// @version 1.0
// @description This is a sample order API
// @host localhost:3001
// @BasePath /
func main() {
	server := gin.Default()
	godotenv.Load()
	db.Connect()
	// Swagger documentation route
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// @Summary Health check endpoint
	// @Description Check if the server is running
	// @Tags health
	// @Produce json
	// @Success 200 {object} map[string]string
	// @Router /health [get]
	server.GET("/health", HealthCheck)

	routes.AuthRoutes(server)

	server.Run(os.Getenv("PORT"))
}

// HealthCheck godoc
// @Summary Health check endpoint
// @Description Check if the server is running
// @Tags health
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func HealthCheck(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "server is fine",
	})
}
