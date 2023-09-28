package user

import (
	"github.com/gin-gonic/gin"
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

// NameIsExist 账户是否已经存在
func NameIsExist(name string) int64 {
	var count int64
	database.DB.
		Where("phone = ?", name).
		Or("email = ?", name).
		Or("name = ?", name).
		Count(&count)
	return count
}

// Get 通过 ID 获取用户
func Get(idstr string) (userModel User) {
	database.DB.Where("id", idstr).First(&userModel)
	return
}

// GetByEmail 通过 Email 来获取用户
func GetByEmail(email string) (userModel User) {
	database.DB.Where("email = ?", email).First(&userModel)
	return
}

// FirstById 通过id获取详细
func FirstById(id uint64) (user User) {
	database.DB.First(&user, id)
	return
}

// FirstPreloadById 通过主键获取详细并加载关联
func FirstPreloadById(id uint64) (user User) {
	database.DB.Preload("Role").First(&user, id)
	return
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
