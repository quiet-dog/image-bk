package initialize

import (
	"image-bk/config"
	"image-bk/global"

	"github.com/qiniu/go-sdk/v7/auth"
)

func InitQiNiu() {
	accessKey := "your access key"
	secretKey := "your secret key"
	if accessKey == "" && global.CONFIG.QiNiu.AccessKey == "" {
		panic("qiniu access key is required")
	}

	var mac *auth.Credentials
	if global.CONFIG.QiNiu.AccessKey != "" {
		mac = auth.New(global.CONFIG.QiNiu.AccessKey, global.CONFIG.QiNiu.SecretKey)
	} else {
		mac = auth.New(accessKey, secretKey)
	}
	global.QiNiu = &config.QiNiuClient{}
	global.QiNiu.Mac = mac
}
