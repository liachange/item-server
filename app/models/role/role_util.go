package role

import (
	"gorm.io/gorm"
	"item-server/app/models"
	"item-server/pkg/app"
	"item-server/pkg/database"
	"item-server/pkg/filter"
	"item-server/pkg/helpers"
	"item-server/pkg/paginator"

	"github.com/gin-gonic/gin"
)

// FindById 通过id获取详细
func FindById(id uint64) (role Role) {
	database.DB.First(&role, id)
	return
}

// FindPreloadById 通过主键获取详细并加载关联
func FindPreloadById(id uint64) (role Role) {
	database.DB.Preload("Permissions", func(tx *gorm.DB) *gorm.DB {
		return tx.Select("id,name,title,guard_name")
	}).First(&role, id)
	return
}

// GetAll 全部角色
func GetAll() (role []Role) {
	database.DB.Select("id", "title", "name").Where("state=?", models.ConstShow()).Find(&role)
	return
}

func Paginate(c *gin.Context, perPage int, filters any) (roles []Role, paging paginator.Paging) {
	query := database.DB.Model(Role{})
	// 拼接查询语句
	filter.QueryFilter(query, helpers.ReqFilter(filters))
	paging = paginator.Paginate(
		c,
		query,
		&roles,
		app.V1URL(database.TableName(&Role{})),
		perPage,
	)
	return
}
