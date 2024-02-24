package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// Logger middleware
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// before request

		c.Next()

		// after request
		latency := time.Since(t)
		log.Print("latency: ", latency)
	}
}
