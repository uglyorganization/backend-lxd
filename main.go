package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	Start()
}

// Start with a /health
func Start() {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {})
	err := r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		panic("Error occurred while running the server: " + err.Error())
	}
}
