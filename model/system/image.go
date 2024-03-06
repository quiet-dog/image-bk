package system

import "gorm.io/gorm"

type ImageModel struct {
	gorm.Model
	Key string `json:"key" gorm:"comment:图片key"`
	//	内容
	Content         string                 `json:"content" gorm:"comment:图片内容"`
	Classifications []*ClassificationModel `json:"classifications" gorm:"many2many:image_classification;"`
	UserModelID     uint                   `json:"userModelID" gorm:"comment:用户id"`
	UserModel       *UserModel             `json:"userModels"`
}

func (ImageModel) TableName() string {
	return "sys_image"
}
