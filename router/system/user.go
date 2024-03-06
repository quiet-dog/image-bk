package system

import (
	"image-bk/api"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (c *UserRouter) InitUserRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	userRouter := Router.Group("user")
	userApi := api.ApiGroupApp.SystemApiGroup.UserApi

	{
		userRouter.POST("login", userApi.Login)
	}
	return userRouter
}
