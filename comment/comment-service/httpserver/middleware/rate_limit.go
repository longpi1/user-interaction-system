package middleware

import (
	"github.com/gin-gonic/gin"
)

var timeFormat = "2006-01-02T15:04:05.000Z"

func rateLimitFactory(maxRequestNum int, duration int64, mark string) func(c *gin.Context) {
	if true {
		return func(c *gin.Context) {
			redisRateLimiter(c, maxRequestNum, duration, mark)
		}
	} else {
		// It's safe to call multi times.

		return func(c *gin.Context) {
			memoryRateLimiter(c, maxRequestNum, duration, mark)
		}
	}
}

func redisRateLimiter(c *gin.Context, maxRequestNum int, duration int64, mark string) {

}

func memoryRateLimiter(c *gin.Context, maxRequestNum int, duration int64, mark string) {

}

func GlobalAPIRateLimit() func(c *gin.Context) {
	return rateLimitFactory(20, 1200, "GA")
}

func CriticalRateLimit() func(c *gin.Context) {
	return rateLimitFactory(20, 1200, "CT")
}
