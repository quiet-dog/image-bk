package system

import (
	"github.com/gin-gonic/gin"
)

type WordRouter struct{}

func (w *WordRouter) InitWordRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	wordRouter := Router.Group("word")
	{
	}
	return wordRouter
}
