package initialize

import (
	"image-bk/global"

	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram/config"
)

func InitWechat() {
	wc := wechat.NewWechat()
	memory := cache.NewMemory()

	global.WX = wc.GetMiniProgram(&config.Config{
		Cache: memory,
	})
}
