package system

import "gorm.io/gorm"

type ClassificationModel struct {
	gorm.Model
	// 名称
	Name   string        `json:"name" gorm:"comment:分类名称"`
	Images []*ImageModel `json:"images" gorm:"many2many:image_classification;"`
}

func (ClassificationModel) TableName() string {
	return "sys_classification"
}
