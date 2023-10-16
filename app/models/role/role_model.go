// Package role 模型 	//Permissions        []*permission.Permission `gorm:"many2many:role_has_permissions" json:"permissions"`
package role

import (
	"errors"
	"gorm.io/gorm"
	"item-server/app/models"
	"item-server/app/models/permission"
	"item-server/app/models/role_has_permission"
	"item-server/pkg/database"
	"item-server/pkg/helpers"
)

type Role struct {
	models.BaseModel
	State       uint8                   `json:"state,omitempty"`
	Name        string                  `json:"name,omitempty"`
	GuardName   string                  `json:"guard,omitempty"`
	Title       string                  `json:"title,omitempty"`
	Description string                  `json:"description,omitempty"`
	Permissions []permission.Permission `gorm:"many2many:role_has_permissions" json:"permissions"`
	models.CommonTimestampsField
}

func (role *Role) Create() {
	database.DB.Create(&role)
}

func (role *Role) CreateMany(permKey []uint64) (id uint64) {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := (tx.Create(&role)).Error; err != nil {
			return err
		}
		var roleHasPermissions []*role_has_permission.RoleHasPermission
		if role.ID > 0 {
			for _, v := range permKey {
				row := &role_has_permission.RoleHasPermission{
					PermissionID: v,
					RoleID:       role.ID,
				}
				roleHasPermissions = append(roleHasPermissions, row)
			}
			if err := (tx.Create(&roleHasPermissions)).Error; err != nil {
				return err
			}
		} else {
			return errors.New("参数异常...")
		}
		return nil
	})
	if err == nil {
		id = role.ID
	}
	return id
}

func (role *Role) Save() (rowsAffected int64) {
	result := database.DB.Save(&role)
	return result.RowsAffected
}

func (role *Role) SaveMany(fieldSelect any, permKey []uint64) (row int) {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := (tx.Select(helpers.ReqSelect(fieldSelect)).Save(&role)).Error; err != nil {
			return err
		}
		if err := (tx.Where("role_id=?", role.ID).Delete(&role_has_permission.RoleHasPermission{})).Error; err != nil {
			return err
		}
		var roleHasPermissions []*role_has_permission.RoleHasPermission
		if role.ID > 0 {
			for _, v := range permKey {
				row := &role_has_permission.RoleHasPermission{
					PermissionID: v,
					RoleID:       role.ID,
				}
				roleHasPermissions = append(roleHasPermissions, row)
			}
			if err := (tx.Create(&roleHasPermissions)).Error; err != nil {
				return err
			}
		} else {
			return errors.New("参数异常...")
		}
		return nil
	})
	if err != nil {
		return 0
	} else {
		return 1
	}
}

func (role *Role) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&role)
	return result.RowsAffected
}
