package system

import (
	"image-bk/global"

	"github.com/qiniu/go-sdk/v7/storage"
)

type QiNiuService struct{}

func (q *QiNiuService) GetToken() (upToken string, err error) {
	putPolicy := storage.PutPolicy{
		Scope: global.CONFIG.QiNiu.Bucket,
	}
	putPolicy.Expires = uint64(global.CONFIG.QiNiu.ExpiresTime)
	upToken = putPolicy.UploadToken(global.QiNiu.Mac)
	return
}
