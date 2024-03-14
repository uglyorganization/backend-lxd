package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware performs api token validation for requests
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("API-Key")
		// Check if API key is empty or invalid
		if apiKey == "" || !isValidAPIKey(apiKey) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		// Proceed with the request
		c.Next()
	}
}

func isValidAPIKey(apiKey string) bool {
	// Logic to validate API key against stored keys
	// You can implement this based on how you manage API keys
	// For demonstration purposes, let's assume a hardcoded API key
	return apiKey == "your-api-key"
}

func main() {
	Start()
}

// Start with a /health
func Start() {
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/health", func(c *gin.Context) {})

	v1 := r.Group("/v1")
	v1.Use(AuthMiddleware())
	{
		v1.GET("/protected", func(c *gin.Context) {})
	}

	err := r.Run(":8080")
	if err != nil {
		panic("Error occurred while running the server: " + err.Error())
	}
}
