package unit

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"item-server/app/replaces"
	"item-server/pkg/app"
	"item-server/pkg/database"
	"item-server/pkg/filter"
	"item-server/pkg/helpers"
	"item-server/pkg/paginator"
)

// FindById 通过id获取详细
func FindById(id uint64) (unit Unit) {
	database.DB.First(&unit, id)
	return
}

// FindPreloadById 通过主键获取详细并加载关联
func FindPreloadById(id uint64) (unit Unit) {
	database.DB.Preload("Category").First(&unit, id)
	return
}

func ScopeCategory(ids []uint64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Joins("inner join category_units as cn on units.id=cn.unit_id").Where("cn.category_id in ?", ids)
	}
}

func Paginate(c *gin.Context, perPage int, filters *replaces.UnitIndex) (units []Unit, paging paginator.Paging) {
	query := database.DB.Model(Unit{})
	if !helpers.Empty(filters.Category) {
		query.Scopes(ScopeCategory(filters.Category))
	}
	// 拼接查询语句
	filter.QueryFilter(query, helpers.ReqFilter(filters))
	paging = paginator.Paginate(
		c,
		query,
		&units,
		app.V1URL(database.TableName(&Unit{})),
		perPage,
	)
	return
}
