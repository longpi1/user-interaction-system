package httpserver

import (
	"github.com/gin-gonic/gin"
	"user-interaction-system/httpserver/controller"
	"user-interaction-system/httpserver/middleware"
	"user-interaction-system/libary/log"
)

func SetRouter(port string) {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.CORS())
	// 设置基础的api相关的由
	setBaseRouter(router)
	// 设置评论相关的路由
	setCommentRouter(router)

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

func setCommentRouter(router *gin.Engine) {
	groupRouter := router.Group("/v1/comment")
	// 认证
	groupRouter.Use(middleware.Auth())
	{
		groupRouter.Use()
		{
			//todo 评论相关操作实现
			groupRouter.POST("/add", controller.AddComment)
			groupRouter.POST("/delete", middleware.AdminAuth(), controller.DeleteComment)
			groupRouter.POST("/list", controller.CommentList)
			groupRouter.POST("/detail", controller.CommentDetail)
		}
	}

}