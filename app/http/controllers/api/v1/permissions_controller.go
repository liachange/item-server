package v1

import (
	"fmt"
	"github.com/spf13/cast"
	"item-server/app/http/resources"
	"item-server/app/models/permission"
	"item-server/app/requests"
	"item-server/pkg/helpers"
	"item-server/pkg/response"

	"github.com/gin-gonic/gin"
)

type PermissionsController struct {
	BaseAPIController
}

func (ctrl *PermissionsController) Index(c *gin.Context) {
	requestFilter := requests.PermissionFilterRequest{}
	if ok := requests.Validate(c, &requestFilter, requests.PermissionFilter); !ok {
		return
	}
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := permission.Paginate(c, 10, &requestFilter)

	response.JSON(c, gin.H{
		"data":  resources.PermissionIndexResource(data),
		"pager": pager,
	})
}

func (ctrl *PermissionsController) Show(c *gin.Context) {
	permissionModel := permission.FindById(cast.ToUint64(c.Param("id")))
	if permissionModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, resources.PermissionShowResource(permissionModel))
}

func (ctrl *PermissionsController) Store(c *gin.Context) {

	request := requests.PermissionRequest{}
	if ok := requests.Validate(c, &request, requests.PermissionCreate); !ok {
		return
	}

	permissionModel := permission.Permission{
		State:       request.State,
		Type:        request.Type,
		Name:        request.Name,
		Title:       request.Title,
		GuardName:   request.Guard,
		Description: request.Description,
		ParentID:    request.Parent,
		Icon:        request.Icon,
		Sort:        request.Sort,
	}
	permissionModel.Create()
	if permissionModel.ID > 0 {
		response.Created(c, permissionModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *PermissionsController) Update(c *gin.Context) {

	permissionModel := permission.FindById(cast.ToUint64(c.Param("id")))
	if permissionModel.ID == 0 {
		response.Abort404(c)
		return
	}
	fmt.Println(c)
	request := requests.PermissionRequest{}

	if ok := requests.Validate(c, &request, requests.PermissionSave); !ok {
		return
	}

	permissionModel.State = request.State
	permissionModel.Type = request.Type
	permissionModel.Name = request.Name
	permissionModel.Title = request.Title
	permissionModel.GuardName = request.Guard
	permissionModel.Description = request.Description
	permissionModel.ParentID = request.Parent
	permissionModel.Icon = request.Icon
	permissionModel.Sort = request.Sort
	permissionModel.UpdatedAt = helpers.TimeNow()
	rowsAffected := permissionModel.Save(&request)
	if rowsAffected > 0 {
		response.Data(c, permissionModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *PermissionsController) Delete(c *gin.Context) {

	permissionModel := permission.FindById(cast.ToUint64(c.Param("id")))
	if permissionModel.ID == 0 {
		response.Abort404(c)
		return
	}

	rowsAffected := permissionModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}
