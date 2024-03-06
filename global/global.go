package global

import (
	"image-bk/config"

	"github.com/go-redis/redis/v8"
	"github.com/silenceper/wechat/v2/miniprogram"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	RUNPATH            = "."
	CONFIG             config.Server
	LOG                *logrus.Logger
	VP                 *viper.Viper
	DB                 *gorm.DB
	WX                 *miniprogram.MiniProgram
	REDIS              *redis.Client
	ConcurrencyControl = &singleflight.Group{}
)
