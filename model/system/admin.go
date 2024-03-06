package system

import "gorm.io/gorm"

type AdminModel struct {
	gorm.Model
	Username string `json:"username" gorm:"comment:用户名"`
	Password string `json:"password" gorm:"comment:密码"`
}

func (AdminModel) TableName() string {
	return "sys_admin"
}
