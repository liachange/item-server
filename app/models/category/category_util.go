package category

import (
	"github.com/gin-gonic/gin"
	"item-server/app/replaces"
	"item-server/pkg/app"
	"item-server/pkg/database"
	"item-server/pkg/filter"
	"item-server/pkg/helpers"
	"item-server/pkg/paginator"
)

// FindById 通过id获取详细
func FindById(id uint64) (category Category) {
	database.DB.First(&category, id)
	return
}
func KeyPluck(key []uint64) (Keys []uint64) {
	database.DB.Model(&Category{}).Where(key).Pluck("id", &Keys)
	return
}

func TreeCategoryAll() (category []*Category) {
	database.DB.Select("id", "title", "sort", "parent_id").Order("level_tree asc").Find(&category)
	return
}

func Paginate(c *gin.Context, perPage int, filters *replaces.CategoryIndex) (categories []Category, paging paginator.Paging) {
	query := database.DB.Model(Category{})
	// 拼接查询语句
	filter.QueryFilter(query, helpers.ReqFilter(filters))
	paging = paginator.Paginate(
		c,
		query,
		&categories,
		app.V1URL(database.TableName(&Category{})),
		perPage,
	)
	return
}
