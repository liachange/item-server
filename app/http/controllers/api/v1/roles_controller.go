package v1

import (
	"github.com/spf13/cast"
	"item-server/app/http/resources"
	"item-server/app/models"
	"item-server/app/models/permission"
	"item-server/app/models/role"
	"item-server/app/requests"
	"item-server/pkg/helpers"
	optimusPkg "item-server/pkg/optimus"
	"item-server/pkg/response"

	"github.com/gin-gonic/gin"
)

type RolesController struct {
	BaseAPIController
}

func (ctrl *RolesController) InitialValue(c *gin.Context) {
	r := resources.Role{PerSlice: permission.GetAll()}
	data := map[string]any{
		"permissions": r.InitialPerSliceResource(),
		"states":      models.InitState(),
	}

	response.Data(c, data)
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
	r := resources.Role{ModelSlice: data}
	response.JSON(c, gin.H{
		"data":  r.IndexResource(),
		"pager": pager,
	})
}

func (ctrl *RolesController) Show(c *gin.Context) {
	id := cast.ToUint64(c.Param("id"))
	if ok := helpers.IdVerify(id); !ok {
		response.Abort404(c)
		return
	}
	roleModel := role.FindPreloadById(optimusPkg.NewOptimus().Decode(id))
	if roleModel.ID == 0 {
		response.Abort404(c)
		return
	}
	r := resources.Role{Model: &roleModel}
	response.Data(c, r.ShowResource())
}

func (ctrl *RolesController) Store(c *gin.Context) {

	request := requests.RoleRequest{}
	if ok := requests.Validate(c, &request, requests.RoleCreate); !ok {
		return
	}
	//权限标识参数处理
	opt := optimusPkg.NewOptimus()
	keys := make([]uint64, 0)
	for _, v := range request.Permission {
		keys = append(keys, opt.Decode(v))
	}
	perKey := permission.KeyPluck(keys)
	roleModel := role.Role{
		State:       request.State,
		Name:        request.Name,
		Title:       request.Title,
		GuardName:   request.Guard,
		Description: request.Description,
	}
	id := roleModel.CreateMany(perKey)
	if id > 0 {
		model := role.FindPreloadById(id)
		r := resources.Role{Model: &model}
		response.Created(c, r.ShowResource())
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *RolesController) Update(c *gin.Context) {
	id := cast.ToUint64(c.Param("id"))
	if ok := helpers.IdVerify(id); !ok {
		response.Abort404(c)
		return
	}
	roleModel := role.FindById(optimusPkg.NewOptimus().Decode(id))
	if roleModel.ID == 0 {
		response.Abort404(c)
		return
	}

	request := requests.RoleRequest{}
	if ok := requests.Validate(c, &request, requests.RoleSave); !ok {
		return
	}
	//权限标识参数处理
	opt := optimusPkg.NewOptimus()
	keys := make([]uint64, 0)
	for _, v := range request.Permission {
		keys = append(keys, opt.Decode(v))
	}
	perKey := permission.KeyPluck(keys)

	roleModel.State = request.State
	roleModel.Name = request.Name
	roleModel.Title = request.Title
	roleModel.GuardName = request.Guard
	roleModel.Description = request.Description
	roleModel.UpdatedAt = helpers.TimeNow()
	rowsAffected := roleModel.SaveMany(&request, perKey)
	if rowsAffected > 0 {
		model := role.FindPreloadById(roleModel.ID)
		r := resources.Role{Model: &model}
		response.Data(c, r.ShowResource())
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *RolesController) Delete(c *gin.Context) {
	id := cast.ToUint64(c.Param("id"))
	if ok := helpers.IdVerify(id); !ok {
		response.Abort404(c)
		return
	}
	roleModel := role.FindById(optimusPkg.NewOptimus().Decode(id))
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
