package system

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Username string `json:"username" gorm:"comment:用户名"`
}

func (UserModel) TableName() string {
	return "sys_user"
}
