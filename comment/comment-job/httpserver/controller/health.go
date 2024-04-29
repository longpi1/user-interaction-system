package controller

import (
	"comment-job/libary/utils"

	"github.com/gin-gonic/gin"
	"github.com/troian/healthcheck"
)

func Ready(c *gin.Context) {
	readyEndpoint := healthcheck.NewHandler().ReadyEndpoint
	utils.RespData(c, "检查是否就绪", readyEndpoint)
}

func Live(c *gin.Context) {
	liveEndpoint := healthcheck.NewHandler().LiveEndpoint
	utils.RespData(c, "检查是否存活", liveEndpoint)
}
