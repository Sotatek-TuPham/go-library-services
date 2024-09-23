package main

import (
	db "library-server/DB"
	"library-server/routes"
	"os"

	_ "library-server/docs" // Import swagger docs

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Library API
// @version 1.0
// @description This is a sample library API
// @host localhost:3000
// @BasePath /
func main() {
	server := gin.Default()
	godotenv.Load()
	db.Connect()

	// Swagger documentation route
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register the health check endpoint
	server.GET("/health", HealthCheck)

	authRoutes := server.Group("/auth")
	routes.AuthRoutes(authRoutes)

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
