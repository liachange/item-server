// Package permission 模型
package permission

import (
	"item-server/app/models"
	"item-server/pkg/database"
)

type Permission struct {
	models.BaseModel

	State       string `json:"state,omitempty"`
	Type        string `json:"type,omitempty"`
	Name        string `json:"name,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Icon        string `json:"icon,omitempty"`
	Sort        string `gorm:"default:1" json:"sort,omitempty"`
	ParentID    string `gorm:"default:0" json:"parent,omitempty"`
	GuardName   string `json:"guard,omitempty"`

	models.CommonTimestampsField
}

func (permission *Permission) Create() {
	database.DB.Create(&permission)
}

func (permission *Permission) Save() (rowsAffected int64) {
	result := database.DB.Save(&permission)
	return result.RowsAffected
}

func (permission *Permission) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&permission)
	return result.RowsAffected
}
