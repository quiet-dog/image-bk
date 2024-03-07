package system

import (
	"image-bk/service"
)

type ApiGroup struct {
	UserApi
}

var (
	userService  = service.ServiceGroupApp.SystemServiceGroup.UserService
	qiNiuService = service.ServiceGroupApp.SystemServiceGroup.QiNiuService
)
