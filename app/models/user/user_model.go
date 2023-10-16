package user

import (
	"errors"
	"gorm.io/gorm"
	"item-server/app/models"
	"item-server/app/models/model_has_role"
	"item-server/app/models/role"
	"item-server/pkg/database"
	"item-server/pkg/hash"
	"item-server/pkg/helpers"
)

// User 用户模型
type User struct {
	models.BaseModel

	State    uint8  `json:"state,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	Avatar   string      `json:"avatar_icon,omitempty"`
	Nickname string      `json:"nickname,omitempty"`
	Role     []role.Role `gorm:"many2many:model_has_roles;joinForeignKey:model_id" json:"roles"`

	models.CommonTimestampsField
}

// Create 创建用户，通过 User.ID 来判断是否创建成功
func (userModel *User) Create() {
	database.DB.Create(&userModel)
}

// ComparePassword 密码是否正确
func (userModel *User) ComparePassword(_password string) bool {
	return hash.BcryptCheck(_password, userModel.Password)
}

// Save 修改信息
func (userModel *User) Save() (rowsAffected int64) {
	result := database.DB.Omit("created_at").Save(&userModel)
	return result.RowsAffected
}

// CreateMany 创建用户并关联角色
func (userModel *User) CreateMany(permKey []uint64) (id uint64) {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := (tx.Create(&userModel)).Error; err != nil {
			return err
		}
		var modelHasRoles []*model_has_role.ModelHasRole
		if userModel.ID > 0 {
			for _, v := range permKey {
				row := &model_has_role.ModelHasRole{
					ModelID:   userModel.ID,
					RoleID:    v,
					ModelType: "user",
				}
				modelHasRoles = append(modelHasRoles, row)
			}
			if err := (tx.Create(&modelHasRoles)).Error; err != nil {
				return err
			}
		} else {
			return errors.New("参数异常...")
		}
		return nil
	})
	if err == nil {
		id = userModel.ID
	}
	return id
}

// SaveMany 修改用户信息和角色关联关系
func (userModel *User) SaveMany(fieldSelect interface{}, permKey []uint64) (rowsAffected int64) {

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		//更新主表信息
		if err := (tx.Select(helpers.ReqSelect(fieldSelect)).Save(&userModel)).Error; err != nil {
			return err
		}
		//删除关联关系
		if err := (tx.Where("model_type=?", "user").Where("model_id=?", userModel.ID).Delete(&model_has_role.ModelHasRole{})).Error; err != nil {
			return err
		}
		var modelHasRoles []*model_has_role.ModelHasRole
		if userModel.ID > 0 {
			for _, v := range permKey {
				row := &model_has_role.ModelHasRole{
					ModelID:   userModel.ID,
					RoleID:    v,
					ModelType: "user",
				}
				modelHasRoles = append(modelHasRoles, row)
			}
			//创建关联关系
			if err := (tx.Create(&modelHasRoles)).Error; err != nil {
				return err
			}
		} else {
			return errors.New("参数异常...")
		}
		return nil
	})
	if err == nil {
		rowsAffected = 1
	}
	return
}

// Delete 删除用户
func (userModel *User) Delete() (rowsAffected int64) {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := (tx.Delete(&userModel)).Error; err != nil {
			return err
		}
		if err := (tx.Where("model_type=?", "user").Where("model_id=?", userModel.ID).Delete(&model_has_role.ModelHasRole{})).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return 0
	} else {
		return 1
	}
}
