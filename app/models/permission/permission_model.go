// Package permission 模型
package permission

import (
	"item-server/app/models"
	"item-server/pkg/database"
	"item-server/pkg/helpers"
)

type Permission struct {
	models.BaseModel

	State       uint8  `json:"state,omitempty"`
	Type        uint8  `json:"type,omitempty"`
	Name        string `json:"name,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Icon        string `json:"icon,omitempty"`
	Sort        uint64 `gorm:"default:1" json:"sort,omitempty"`
	ParentID    uint64 `gorm:"default:0" json:"parent,omitempty"`
	GuardName   string `json:"guard,omitempty"`

	models.CommonTimestampsField
}

func (permission *Permission) Create() {
	database.DB.Create(&permission)
}

func (permission *Permission) Save(fieldSelect interface{}) (rowsAffected int64) {
	result := database.DB.Select(helpers.ReqSelect(fieldSelect)).Save(&permission)
	return result.RowsAffected
}

func (permission *Permission) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&permission)
	return result.RowsAffected
}
