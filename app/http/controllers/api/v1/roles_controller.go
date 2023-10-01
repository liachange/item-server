package v1

import (
	"github.com/spf13/cast"
	"item-server/app/http/resources"
	"item-server/app/models/permission"
	"item-server/app/models/role"
	"item-server/app/requests"
	"item-server/pkg/helpers"
	"item-server/pkg/response"

	"github.com/gin-gonic/gin"
)

type RolesController struct {
	BaseAPIController
}

func (ctrl *RolesController) Index(c *gin.Context) {
	requestFilter := requests.RoleFilterRequest{}
	if ok := requests.Validate(c, &requestFilter, requests.RoleFilter); !ok {
		return
	}
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := role.Paginate(c, 10, &requestFilter)
	response.JSON(c, gin.H{
		"data":  resources.RoleIndexResource(data),
		"pager": pager,
	})
}

func (ctrl *RolesController) Show(c *gin.Context) {
	roleModel := role.FindPreloadById(cast.ToUint64(c.Param("id")))
	if roleModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, resources.RoleShowResource(roleModel))
}

func (ctrl *RolesController) Store(c *gin.Context) {

	request := requests.RoleRequest{}
	if ok := requests.Validate(c, &request, requests.RoleCreate); !ok {
		return
	}
	perKey := permission.KeyPluck(request.Permission)
	roleModel := role.Role{
		State:       request.State,
		Name:        request.Name,
		Title:       request.Title,
		GuardName:   request.Guard,
		Description: request.Description,
	}
	rowsAffected := roleModel.CreateMany(perKey)
	if rowsAffected > 0 {
		response.Created(c, roleModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *RolesController) Update(c *gin.Context) {

	roleModel := role.FindById(cast.ToUint64(c.Param("id")))
	if roleModel.ID == 0 {
		response.Abort404(c)
		return
	}

	request := requests.RoleRequest{}
	if ok := requests.Validate(c, &request, requests.RoleSave); !ok {
		return
	}
	perKey := permission.KeyPluck(request.Permission)

	roleModel.State = request.State
	roleModel.Name = request.Name
	roleModel.Title = request.Title
	roleModel.GuardName = request.Guard
	roleModel.Description = request.Description
	roleModel.UpdatedAt = helpers.TimeNow()
	rowsAffected := roleModel.SaveMany(&request, perKey)
	if rowsAffected > 0 {
		response.Data(c, roleModel)
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *RolesController) Delete(c *gin.Context) {

	roleModel := role.FindById(cast.ToUint64(c.Param("id")))
	if roleModel.ID == 0 {
		response.Abort404(c)
		return
	}

	rowsAffected := roleModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}
