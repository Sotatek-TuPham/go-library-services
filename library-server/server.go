package main

import (
	db "library-server/DB"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	server := gin.Default()
	godotenv.Load()
	db.Connect()
	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "server is fine",
		})
	})

	server.Run(os.Getenv("PORT"))
}
