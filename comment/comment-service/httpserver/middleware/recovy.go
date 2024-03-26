package middleware

import (
	"comment-service/libary/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 自定义Recovery中间件
func CustomRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 记录错误信息
				log.Error("panic: %v\n", err)
				// 发送自定义错误响应
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code":    http.StatusInternalServerError,
					"message": "Internal Server Error",
				})
			}
		}()
		c.Next() // 处理下一个中间件或路由
	}
}
