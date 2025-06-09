package handlers

import (
	"log"
	"net/http"
	"weather-api/cache"
	"weather-api/weather"

	"github.com/gin-gonic/gin"
)

type WeatherHandler struct {
	Service weather.WeatherService
	Cache   cache.RedisCache
}

func NewWeatherHandler(service weather.WeatherService, cache cache.RedisCache) *WeatherHandler {
	return &WeatherHandler{
		Service: service,
		Cache:   cache,
	}
}

func (h *WeatherHandler) GetWeatherByCity(c *gin.Context) {
	city := c.Query("city")
	if city == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "city parameter is required"})
		return
	}

	// Check in cache first
	if data, found := h.Cache.Get(city); found {
		log.Println("✅ Cache hit for:", city)
		c.JSON(http.StatusOK, data)
		return
	}

	log.Println("❌ Cache miss for:", city)

	// Fetch from weather service
	weatherData, err := h.Service.GetCurrentWeather(city)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Save to cache
	h.Cache.Set(city, weatherData)

	c.JSON(http.StatusOK, weatherData)
}
