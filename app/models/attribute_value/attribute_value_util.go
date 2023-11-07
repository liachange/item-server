package attribute_value

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
func FindById(id uint64) (attributeValue AttributeValue) {
	database.DB.First(&attributeValue, id)
	return
}

// FindPreloadById 通过主键获取详细并加载关联
func FindPreloadById(id uint64) (attributeValue AttributeValue) {
	database.DB.Preload("AttributeName").First(&attributeValue, id)
	return
}

func Paginate(c *gin.Context, perPage int, rep *replaces.AttributeValueIndex) (attributeValues []*AttributeValue, paging paginator.Paging) {
	query := database.DB.Model(AttributeValue{})

	// 拼接查询语句
	filter.QueryFilter(query, helpers.ReqFilter(rep))
	paging = paginator.Paginate(
		c,
		query,
		&attributeValues,
		app.V1URL(database.TableName(&AttributeValue{})),
		perPage,
	)
	return
}
