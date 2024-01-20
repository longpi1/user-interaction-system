package httpserver

import (
	"github.com/gin-gonic/gin"
	"model-api/httpserver/controller"
	"model-api/httpserver/middleware"
	"model-api/libary/log"
)

func SetRouter() {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.CORS())
	// 设置api相关的路由
	setApiRouter(router)
	// 设置大模型相关的路由
	setModelRouter(router)
	err := router.Run(":" + "8888")
	if err != nil {
		log.Error("server run fail")
	}
}

func setApiRouter(router *gin.Engine) {
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
		userRouter := apiRouter.Group("/user")
		{
			userRouter.GET("register", controller.Register)
			userRouter.POST("/login", middleware.CriticalRateLimit(), controller.Login)
			userRouter.GET("/logout", controller.Logout)
			//todo 获取用户信息，设置token信息等
		}
		logRouter := apiRouter.Group("/log")
		logRouter.GET("/", middleware.AdminAuth(), controller.GetLogs)
		logRouter.DELETE("/delete", middleware.AdminAuth(), controller.DeleteLogs)
		groupRoute := apiRouter.Group("/group")
		groupRoute.Use(middleware.AdminAuth())
		{
			groupRoute.GET("/")
		}
	}
}

func setModelRouter(router *gin.Engine) {
	groupRouter := router.Group("/v1")
	// 认证
	groupRouter.Use(middleware.Auth())
	{
		groupRouter.Use()
		{
			//todo 模型相关操作实现
			groupRouter.POST("/completions")
			groupRouter.POST("/chat/completions")
			groupRouter.POST("/edits")
			groupRouter.POST("/images/generations")
			groupRouter.POST("/images/edits")
			groupRouter.POST("/images/variations")
			groupRouter.POST("/embeddings")
			groupRouter.POST("/engines/:model/embeddings")
			groupRouter.POST("/audio/transcriptions")
			groupRouter.POST("/audio/translations")
			groupRouter.POST("/audio/speech")
			groupRouter.GET("/files")
			groupRouter.POST("/files")
			groupRouter.DELETE("/files/:id")
			groupRouter.GET("/files/:id")
			groupRouter.GET("/files/:id/content")
			groupRouter.POST("/fine_tuning/jobs")
			groupRouter.GET("/fine_tuning/jobs")
			groupRouter.GET("/fine_tuning/jobs/:id")
			groupRouter.POST("/fine_tuning/jobs/:id/cancel")
			groupRouter.GET("/fine_tuning/jobs/:id/events")
		}
		modelRouter := groupRouter.Group("/models")
		{
			modelRouter.GET("")
			modelRouter.GET("/:model")
			modelRouter.DELETE("/models/:model")
		}

	}

}
