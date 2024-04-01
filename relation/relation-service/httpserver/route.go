package httpserver

import (
	"relation-service/httpserver/controller"
	"relation-service/httpserver/middleware"
	"relation-service/libary/log"

	"github.com/gin-gonic/gin"
)

func SetRouter(port string) {
	router := gin.New()
	router.Use(middleware.CORS())
	// 设置Recovery中间件，主要用于拦截paic错误，不至于导致进程崩掉
	router.Use(gin.Recovery())
	// 日志记录耗时
	router.Use(middleware.Logger())
	// 设置基础的api相关的由
	setBaseRouter(router)
	// 设置关注相关的路由
	setRelationRouter(router)

	err := router.Run(":" + port)
	if err != nil {
		log.Error("server run fail")
	}
}

func setBaseRouter(router *gin.Engine) {
	// api相关api
	apiRouter := router.Group("/api")
	// 频率限制
	apiRouter.Use(middleware.GlobalAPIRateLimit())
	{

		introduceRouter := apiRouter.Group("/introduce")
		{
			//todo 网站升级迭代信息查找等
			introduceRouter.GET("")
		}
		oauthRouter := apiRouter.Group("/oauth")
		{
			//todo 通过github或者微信认证登录
			oauthRouter.GET("")
		}
	}
}

func setRelationRouter(router *gin.Engine) {
	groupRouter := router.Group("/v1/relation")
	// 认证
	groupRouter.Use(middleware.Auth())
	{
		groupRouter.Use()
		{
			//todo 评论相关操作实现
			groupRouter.POST("/relation", controller.Relation)
			groupRouter.POST("/relation_count", controller.RelationCount)
			groupRouter.POST("/following", controller.Following)
			groupRouter.POST("/fans", controller.Fans)
			groupRouter.POST("/isFollowing", controller.IsFollowing)
			groupRouter.POST("/isFollowingBatch", controller.IsFollowingBatch)
			groupRouter.POST("/mutualFollow", controller.MutualFollow)
		}
	}
}
