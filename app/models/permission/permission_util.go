package permission

import (
	"item-server/pkg/app"
	"item-server/pkg/database"
	"item-server/pkg/filter"
	"item-server/pkg/helpers"
	"item-server/pkg/paginator"

	"github.com/gin-gonic/gin"
)

// FirstById 通过id获取详细
func FirstById(id uint64) (perm Permission) {
	database.DB.First(&perm, id)
	return
}

func KeyPluck(key []uint64) (Keys []uint64) {
	database.DB.Model(&Permission{}).Where(key).Pluck("id", &Keys)
	return
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
