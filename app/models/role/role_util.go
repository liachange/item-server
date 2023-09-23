package role

import (
	"gorm.io/gorm"
	"item-server/pkg/app"
	"item-server/pkg/database"
	"item-server/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (role Role) {
	database.DB.Preload("Permissions", func(db *gorm.DB) *gorm.DB {
		return db.Select("id,name,title")
	}).First(&role, idstr)
	return
}
func First(id string) (role Role) {
	database.DB.First(&role, id)
	return
}

func GetBy(field, value string) (role Role) {
	database.DB.Where("? = ?", field, value).First(&role)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Role{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (roles []Role, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Role{}),
		&roles,
		app.V1URL(database.TableName(&Role{})),
		perPage,
	)
	return
}
