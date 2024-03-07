package main

import (
	"image-bk/core"
	"image-bk/global"
	"image-bk/initialize"
)

func main() {

	global.VP = core.Viper()
	global.LOG = core.InitLogger()
	// 微信初始化
	initialize.InitWechat()
	// 七牛云初始化
	initialize.InitQiNiu()
	core.RunServer()
}
