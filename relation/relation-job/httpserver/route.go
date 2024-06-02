package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/longpi1/user-interaction-system/relation-job/httpserver/controller"

	"github.com/longpi1/user-interaction-system/relation-job/httpserver/middleware"

	"github.com/longpi1/gopkg/libary/log"

	"github.com/gin-contrib/pprof"
)

func SetRouter(port string) {
	router := gin.New()
	router.Use(middleware.CORS())
	// 设置Recovery中间件，主要用于拦截paic错误，不至于导致进程崩掉
	router.Use(gin.Recovery())
	// 日志记录耗时
	router.Use(middleware.Logger())

	setRouter(router)

	err := router.Run(":" + port)
	if err != nil {
		log.Error("server run fail")
	}
}

func setRouter(router *gin.Engine) {
	// 设置pprof接口用来分析
	pprof.Register(router, "debug/pprof")
	// api相关api
	apiRouter := router.Group("/api")
	// 频率限制
	apiRouter.Use(middleware.GlobalAPIRateLimit())
	{
		apiRouter.GET("/ready", controller.Ready) // GET /ready?full=1
		apiRouter.GET("/live", controller.Live)   // GET /live?full=1
		introduceRouter := apiRouter.Group("/introduce")
		{
			//todo 网站升级迭代信息查找等
			introduceRouter.GET("")
		}
	}
}
