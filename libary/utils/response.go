package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"user-interaction-system/libary/log"
)

func RespData(c *gin.Context, msg string, data interface{}) {
	LogWithHttpInfo(c, false)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": msg,
		"data":    data,
	})
}

func RespSuccess(c *gin.Context, msg string) {
	LogWithHttpInfo(c, false)
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
		"success": true,
	})
}

func RespError(c *gin.Context, msg string) {
	LogWithHttpInfo(c, true)
	c.JSON(http.StatusOK, gin.H{
		"message": msg,
		"success": false,
	})
}

// LogWithHttpInfo 打印http相关参数等日志
func LogWithHttpInfo(c *gin.Context, isError bool) {
	fields := make(map[string]interface{})
	req := c.Request

	path := string(req.Method)
	if path == "" {
		path = "/"
	}
	fields["url"] = req.RequestURI
	fields["params"] = c.Params
	fields["path"] = path
	fields["host"] = req.Host
	fields["status"] = req.Response.Status
	fields["remote_ip"] = req.RemoteAddr
	fields["client_ip"] = c.ClientIP()
	fields["error"] = c.Errors.JSON()
	if isError {
		log.Error(fields)
	} else {
		log.Info(fields)
	}

}
