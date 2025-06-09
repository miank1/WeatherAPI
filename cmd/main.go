package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"weather-api/cache"
	"weather-api/handlers"
	"weather-api/weather"

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

	weatherService := weather.NewOpenWeatherClient(apiKey)
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPwd := os.Getenv("REDIS_PASSWORD")
	cache := cache.NewRedisCache(redisAddr, redisPwd, 10*time.Minute)
	handler := handlers.NewWeatherHandler(weatherService, *cache)
	router.GET("/weather", handler.GetWeatherByCity)

	fmt.Println("Server running on port")
	err = router.Run(":" + port)
	if err != nil {
		log.Fatal("❌ Failed to start server:", err)
	}

}
