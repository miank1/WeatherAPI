package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	apiKey := os.Getenv("WEATHER_API_KEY")
	port := os.Getenv("port")
	if apiKey == "" {
		log.Fatal("WEATHER_API_KEY not set ")
	}
	if port == "" {
		port = "8080"
	}

	fmt.Println("✅ API Key Loaded: ", apiKey)
	fmt.Println("✅ Server will run on port :", port)

	// Initializing the router
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	fmt.Println("Server running on port")
	err = router.Run(":" + port)
	if err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}

}
