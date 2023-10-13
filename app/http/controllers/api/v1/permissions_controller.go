package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"item-server/app/http/resources"
	"item-server/app/models/permission"
	"item-server/app/requests"
	"item-server/pkg/helpers"
	optimusPkg "item-server/pkg/optimus"
	"item-server/pkg/response"
)

type PermissionsController struct {
	BaseAPIController
}

func (ctrl *PermissionsController) InitialValue(c *gin.Context) {
	r := resources.Permission{ModelSlice: permission.PageCategory()}
	data := map[string]any{
		"categories": r.InitialResource(),
		"genres":     permission.InitGenre(),
		"states":     permission.InitState(),
	}

	response.Data(c, data)
}

func (ctrl *PermissionsController) Index(c *gin.Context) {
	var resource resources.Permission
	requestFilter := requests.PermissionFilterRequest{}

	if ok := requests.Validate(c, &requestFilter, requests.PermissionFilter); !ok {
		return
	}
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := permission.Paginate(c, 10, &requestFilter)
	resource.ModelSlice = data
	response.JSON(c, gin.H{
		"data":  resource.IndexResource(),
		"pager": pager,
	})
}

func (ctrl *PermissionsController) Show(c *gin.Context) {
	id := cast.ToUint64(c.Param("id"))
	if ok := helpers.IdVerify(id); !ok {
		return
	}
	permissionModel := permission.FindById(optimusPkg.NewOptimus().Decode(id))
	if permissionModel.ID == 0 {
		response.Abort404(c)
		return
	}
	// 数据封装
	r := resources.Permission{Model: &permissionModel}
	response.Data(c, r.ShowResource())
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
		ParentID:    optimusPkg.NewOptimus().Decode(request.Parent),
		Icon:        request.Icon,
		Sort:        request.Sort,
	}
	permissionModel.Create()
	if permissionModel.ID > 0 {
		model := permission.FindById(permissionModel.ID)
		// 数据封装
		r := resources.Permission{Model: &model}
		response.Created(c, r.ShowResource())
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *PermissionsController) Update(c *gin.Context) {
	id := cast.ToUint64(c.Param("id"))
	fmt.Println(helpers.IdVerify(id))
	if ok := helpers.IdVerify(id); !ok {
		return
	}
	permissionModel := permission.FindById(optimusPkg.NewOptimus().Decode(id))
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
	permissionModel.ParentID = optimusPkg.NewOptimus().Decode(request.Parent)
	permissionModel.Icon = request.Icon
	permissionModel.Sort = request.Sort
	permissionModel.UpdatedAt = helpers.TimeNow()
	rowsAffected := permissionModel.Save(&request)
	if rowsAffected > 0 {
		model := permission.FindById(permissionModel.ID)
		// 数据封装
		r := resources.Permission{Model: &model}
		response.Data(c, r.ShowResource())
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *PermissionsController) Delete(c *gin.Context) {
	id := cast.ToUint64(c.Param("id"))
	if ok := helpers.IdVerify(id); !ok {
		return
	}
	permissionModel := permission.FindById(optimusPkg.NewOptimus().Decode(id))
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
