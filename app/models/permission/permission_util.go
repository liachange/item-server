package permission

import (
    "item-server/pkg/database"
    "item-server/pkg/paginator"
    "item-server/pkg/app"

    "github.com/gin-gonic/gin"
)

func Get(idstr string) (permission Permission) {
    database.DB.Where("id", idstr).First(&permission)
    return
}

func GetBy(field, value string) (permission Permission) {
    database.DB.Where("? = ?", field, value).First(&permission)
    return
}

func All() (permissions []Permission) {
    database.DB.Find(&permissions)
    return
}

func IsExist(field, value string) bool {
    var count int64
    database.DB.Model(Permission{}).Where("? = ?", field, value).Count(&count)
    return count > 0
}

func Paginate(c *gin.Context, perPage int) (permissions []Permission, paging paginator.Paging) {
    paging = paginator.Paginate(
        c,
        database.DB.Model(Permission{}),
        &permissions,
        app.V1URL(database.TableName(&Permission{})),
        perPage,
    )
    return
}