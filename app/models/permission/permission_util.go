package permission

import (
	"item-server/pkg/app"
	"item-server/pkg/database"
	"item-server/pkg/filter"
	"item-server/pkg/helpers"
	"item-server/pkg/paginator"

	"github.com/gin-gonic/gin"
)

const (
	page = iota + 1
	btn
)
const (
	show = iota + 1
	hide
)

// FindById 通过id获取详细
func FindById(id uint64) (perm Permission) {
	database.DB.First(&perm, id)
	return
}

func KeyPluck(key []uint64) (Keys []uint64) {
	database.DB.Model(&Permission{}).Where(key).Pluck("id", &Keys)
	return
}

// PageCategory 页面分类
func PageCategory() (perm []Permission) {
	database.DB.Select("id", "title", "name").Where("parent_id=?", 0).Where("type=?", page).Find(&perm)
	return
}

func InitGenre() []map[string]any {
	return []map[string]any{
		{
			"value": page,
			"label": "页面",
		},
		{
			"value": btn,
			"label": "按钮",
		},
	}
}
func InitState() []map[string]any {
	return []map[string]any{
		{
			"value": show,
			"label": "显示",
		},
		{
			"value": hide,
			"label": "隐藏",
		},
	}
}
func Paginate(c *gin.Context, perPage int, filters interface{}) (permissions []Permission, paging paginator.Paging) {
	query := database.DB.Model(Permission{})
	// 拼接查询语句
	filter.QueryFilter(query, helpers.ReqFilter(filters))
	paging = paginator.Paginate(
		c,
		query,
		&permissions,
		app.V1URL(database.TableName(&Permission{})),
		perPage,
	)
	return
}
