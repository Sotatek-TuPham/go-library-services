package main

import (
	db "library-server/DB"
	"library-server/adapter"
	"library-server/routes"
	"log"
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
	// Initialize RabbitMQ connection
	rabbitMQURL := os.Getenv("RABBITMQ_URL")
	if rabbitMQURL == "" {
		rabbitMQURL = "amqp://guest:guest@localhost:5672/"
	}

	rabbitmq, err := adapter.NewRabbitMQ(rabbitMQURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer rabbitmq.Close()

	log.Println("Connected to RabbitMQ")
	server := gin.Default()
	godotenv.Load()
	db.InitializeDatabase()

	// Swagger documentation route
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register the health check endpoint
	server.GET("/health", HealthCheck)

	authRoutes := server.Group("/auth")
	routes.AuthRoutes(authRoutes)

	bookRoutes := server.Group("/books")
	routes.BookRoutes(bookRoutes)

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
