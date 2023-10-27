package brand

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
func FindById(id uint64) (brand Brand) {
	database.DB.First(&brand, id)
	return
}

// FindPreloadById 通过主键获取详细并加载关联
func FindPreloadById(id uint64) (brand Brand) {
	database.DB.Preload("Category").First(&brand, id)
	return
}

func ScopeCategory(ids []uint64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Joins("inner join category_brands as cb  on brands.id=cb.brand_id").Where("cb.category_id in ?", ids)
	}
}

func Paginate(c *gin.Context, perPage int, rep *replaces.BrandIndex) (brands []Brand, paging paginator.Paging) {
	query := database.DB.Model(Brand{})
	//分类查询
	if !helpers.Empty(rep.Category) {
		query.Scopes(ScopeCategory(rep.Category))
	}
	// 拼接查询语句
	filter.QueryFilter(query, helpers.ReqFilter(rep))
	paging = paginator.Paginate(
		c,
		query,
		&brands,
		app.V1URL(database.TableName(&Brand{})),
		perPage,
	)
	return
}
