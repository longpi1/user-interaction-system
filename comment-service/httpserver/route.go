package httpserver

import (
	"comment-service/httpserver/controller"
	"comment-service/httpserver/middleware"
	"comment-service/libary/log"

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
	// 设置评论相关的路由
	setCommentRouter(router)
	// 设置资源的相关路由
	setResourceRoute(router)

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
			groupRouter.POST("/interact", controller.CommentInteract)
			groupRouter.POST("/top", controller.CommentTop)
			groupRouter.POST("/highlight", controller.CommentHighlight)
			groupRouter.POST("/detail", controller.CommentDetail)
		}
	}
}

func setResourceRoute(router *gin.Engine) {
	resourceRouter := router.Group("/v1/resource")
	// 认证
	resourceRouter.Use(middleware.Auth())
	{
		resourceRouter.POST("/add", middleware.AdminAuth(), controller.AddResource)
		resourceRouter.POST("/delete", middleware.AdminAuth(), controller.DeleteResource)
		resourceRouter.POST("/list", controller.GetResourceList)
		resourceRouter.POST("/detail", controller.ResourceDetail)
	}
}
