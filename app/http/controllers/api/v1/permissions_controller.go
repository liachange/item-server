package v1

import (
	"item-server/app/models/permission"
	"item-server/app/requests"
	"item-server/pkg/response"

	"github.com/gin-gonic/gin"
)

type PermissionsController struct {
	BaseAPIController
}

func (ctrl *PermissionsController) Index(c *gin.Context) {
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := permission.Paginate(c, 10)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}

func (ctrl *PermissionsController) Show(c *gin.Context) {
	permissionModel := permission.Get(c.Param("id"))
	if permissionModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, permissionModel)
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

	permissionModel := permission.Get(c.Param("id"))
	if permissionModel.ID == 0 {
		response.Abort404(c)
		return
	}

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
	rowsAffected := permissionModel.Save()
	if rowsAffected > 0 {
		response.Data(c, permissionModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *PermissionsController) Delete(c *gin.Context) {

	permissionModel := permission.Get(c.Param("id"))
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
