package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/longpi1/gopkg/libary/http/middleware"
	"github.com/longpi1/gopkg/libary/log"
	"github.com/longpi1/user-interaction-system/like-service/httpserver/controller"
)

func SetRouter(port string) {
	router := gin.New()
	router.Use(middleware.CORS())
	// 设置Recovery中间件，主要用于拦截paic错误，不至于导致进程崩掉
	router.Use(gin.Recovery())
	// 日志记录耗时
	router.Use(middleware.Logger())

	// 设置评论相关的路由
	setCommentRouter(router)
	err := router.Run(":" + port)
	if err != nil {
		log.Error("server run fail")
	}
}

func setCommentRouter(router *gin.Engine) {
	groupRouter := router.Group("/v1/like")
	// 认证
	groupRouter.Use(middleware.Auth())
	{
		groupRouter.Use()
		{
			//todo 点赞相关操作实现
			groupRouter.POST("/operation", controller.)
			groupRouter.POST("/list", controller.)
		}
	}
}
