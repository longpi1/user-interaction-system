package httpserver

import "github.com/gin-gonic/gin"

var server *gin.Engine

func GetServer() *gin.Engine {
	if server == nil {
		server = gin.New()
	}
	return server
}
