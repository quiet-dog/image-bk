package initialize

import (
	"image-bk/middleware"
	"image-bk/middleware/log"
	"image-bk/router"

	docs "image-bk/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouters() *gin.Engine {
	docs.SwaggerInfo.BasePath = "/api/v1"
	Router := gin.New()
	Router.Use(log.GinLogger()).Use(middleware.Cors())
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	systemRouter := router.RouterGroupApp.System
	RouterGroup := Router.Group("")
	{
		systemRouter.InitWordRouter(RouterGroup)
	}

	WcRouterGroup := Router.Group("wpwei")
	{
		systemRouter.InitUserRouter(WcRouterGroup)
	}
	return Router
}
