package system

import (
	"image-bk/model/common/request"
	"image-bk/model/common/response"

	"github.com/gin-gonic/gin"
)

type UserApi struct{}

// @Summary		微信小程序登陆
// @Description	微信小程序登陆
// @Tags			小程序
// @Accept			application/json
// @Product		application/json
// @Param			data	body		request.WxLogin	true	"data"
// @Success		200		{string}	string					"{"success":true,"data":{},"msg":"创建成功"}"
// @Router			/mpwei/user/login [post]
func (u *UserApi) Login(c *gin.Context) {
	var req request.WxLogin
	err := c.ShouldBindJSON(&req)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}

	if err := userService.Login(c); err != nil {
		response.FailWithMessage(err.Error(), c)
	} else {
		response.OkWithMessage("登录成功", c)
	}
}
