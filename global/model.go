package global

import (
	"time"

	"gorm.io/gorm"
)

type MODEL struct {
	ID        uint      `gorm:"PRIMARY_KEY;AUTO_INCREMENT" nestedset:"id"` // 主键ID
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
	DeletedAt gorm.DeletedAt
	Updated   int64 `json:"updated" gorm:"autoUpdateTime:milli"`
	Created   int64 `json:"created" gorm:"autoCreateTime:milli"`
	Deleted   int64 `json:"deleted" gorm:""`
}
