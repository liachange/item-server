package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"item-server/app/models"
	"item-server/app/models/role"
	"item-server/pkg/app"
	"item-server/pkg/database"
	"item-server/pkg/filter"
	"item-server/pkg/helpers"
	"item-server/pkg/paginator"
)

// IsEmailExist 判断 Email 已被注册
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// IsPhoneExist 判断手机号已被注册
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}

// GetByPhone 通过手机号来获取用户
func GetByPhone(phone string) (userModel User) {
	database.DB.Where("phone = ?", phone).First(&userModel)
	return
}

// GetByMulti 通过 手机号/Email/用户名 来获取用户
func GetByMulti(loginID string) (userModel User) {
	database.DB.
		Select("id", "password", "nickname").
		Where("phone = ?", loginID).
		Or("email = ?", loginID).
		Or("name = ?", loginID).
		First(&userModel)
	return
}

// Get 通过 ID 获取用户
func Get(idstr string) (userModel User) {
	database.DB.Where("state=?", models.ConstShow()).Where("id", idstr).First(&userModel)
	return
}

// GetById 外部使用
func GetById(id uint64) (userModel User) {
	database.DB.Where("state=?", models.ConstShow()).First(&userModel, id)
	return
}

// GetByEmail 通过 Email 来获取用户
func GetByEmail(email string) (userModel User) {
	database.DB.Where("email = ?", email).First(&userModel)
	return
}

// FindById 通过id获取详细
func FindById(id uint64) (user User) {
	database.DB.First(&user, id)
	return
}

// FindPreloadById 通过主键获取详细并加载关联
func FindPreloadById(id uint64) (user User) {
	database.DB.Preload("Role").First(&user, id)
	return
}

func FindUserMenu(id uint64) (user UserMany) {
	database.DB.Preload("Roles.Permissions", func(db *gorm.DB) *gorm.DB {
		return db.Order("permissions.sort desc")
	}).Find(&user, id)
	return user
}

func KeyPluck(key []uint64) (Keys []uint64) {
	database.DB.Model(&role.Role{}).Where(key).Pluck("id", &Keys)
	return
}

// Paginate 分页
func Paginate(c *gin.Context, perPage int, filters interface{}) (users []User, paging paginator.Paging) {

	query := database.DB.Model(User{})
	// 拼接查询语句
	filter.QueryFilter(query, helpers.ReqFilter(filters))
	paging = paginator.Paginate(
		c,
		query,
		&users,
		app.V1URL(database.TableName(&User{})),
		perPage,
	)
	return
}
