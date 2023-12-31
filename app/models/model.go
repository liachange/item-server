package models

import (
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"time"
)

const (
	show = iota + 1
	hide
)

// BaseModel 模型基类
type BaseModel struct {
	ID uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
}

// CommonTimestampsField 时间戳
type CommonTimestampsField struct {
	CreatedAt time.Time `gorm:"column:created_at;index;" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;index;" json:"updated_at,omitempty"`
}

// DeletedAt 删除字段
type DeletedAt struct {
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index;default:NULL;" json:"deleted_at,omitempty"`
}

// GetStringID 获取 ID 的字符串格式
func (a BaseModel) GetStringID() string {
	return cast.ToString(a.ID)
}

func InitState() []map[string]any {
	return []map[string]any{
		{
			"value": show,
			"label": "显示",
		},
		{
			"value": hide,
			"label": "隐藏",
		},
	}
}
func ConstShow() uint8 {
	return show
}
func ConstHide() uint8 {
	return hide
}
