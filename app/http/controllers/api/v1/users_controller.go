package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"item-server/app/http/resources"
	"item-server/app/models"
	"item-server/app/models/role"
	"item-server/app/models/user"
	"item-server/app/requests"
	"item-server/pkg/auth"
	"item-server/pkg/config"
	"item-server/pkg/file"
	"item-server/pkg/hash"
	"item-server/pkg/helpers"
	optimusPkg "item-server/pkg/optimus"
	"item-server/pkg/response"
)

type UsersController struct {
	BaseAPIController
}

func (ctrl *UsersController) InitialValue(c *gin.Context) {
	r := resources.User{RoleSlice: role.GetAll()}
	data := map[string]any{
		"roles":  r.InitialRoleResource(),
		"states": models.InitState(),
	}

	response.Data(c, data)
}

func (ctrl *UsersController) Index(c *gin.Context) {
	requestFilter := requests.UserFilterRequest{}
	if ok := requests.Validate(c, &requestFilter, requests.UserFilter); !ok {
		return
	}
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := user.Paginate(c, 10, &requestFilter)
	r := resources.User{ModelSlice: data}
	response.JSON(c, gin.H{
		"data":  r.IndexResource(),
		"pager": pager,
	})
}

func (ctrl *UsersController) Show(c *gin.Context) {
	id := cast.ToUint64(c.Param("id"))
	if ok := helpers.IdVerify(id); !ok {
		response.Abort404(c)
		return
	}
	userModel := user.FindPreloadById(optimusPkg.NewOptimus().Decode(id))
	if userModel.ID == 0 {
		response.Abort404(c)
		return
	}
	r := resources.User{Model: &userModel}
	response.Data(c, r.ShowResource())
}

func (ctrl *UsersController) Store(c *gin.Context) {

	request := requests.UserRequest{}
	if ok := requests.Validate(c, &request, requests.UserCreate); !ok {
		return
	}
	//角色标识参数处理
	opt := optimusPkg.NewOptimus()
	ids := make([]uint64, 0)
	for _, v := range request.Role {
		ids = append(ids, opt.Decode(v))
	}
	userModel := user.User{
		Name:     request.Name,
		Password: hash.BcryptHash(request.Password),
		State:    request.State,
		Nickname: request.Nickname,
		Phone:    request.Phone,
		Email:    request.Email,
	}
	id := userModel.CreateMany(ids)
	if id > 0 {
		model := user.FindPreloadById(id)
		r := resources.User{Model: &model}
		response.Created(c, r.ShowResource())
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *UsersController) Update(c *gin.Context) {
	id := cast.ToUint64(c.Param("id"))
	if ok := helpers.IdVerify(id); !ok {
		response.Abort404(c)
		return
	}
	userModel := user.FindById(optimusPkg.NewOptimus().Decode(id))
	if userModel.ID == 0 {
		response.Abort404(c)
		return
	}
	request := requests.UserRequest{}
	if ok := requests.Validate(c, &request, requests.UserSave); !ok {
		return
	}
	//角色标识参数处理
	opt := optimusPkg.NewOptimus()
	ids := make([]uint64, 0)
	for _, v := range request.Role {
		ids = append(ids, opt.Decode(v))
	}

	userModel.State = request.State
	userModel.Name = request.Name
	userModel.UpdatedAt = helpers.TimeNow()
	userModel.Phone = request.Phone
	userModel.Email = request.Email
	userModel.Nickname = request.Nickname

	if len(cast.ToString(request.Password)) > 0 {
		userModel.Password = hash.BcryptHash(request.Password)
	}
	rowsAffected := userModel.SaveMany(&request, ids)
	if rowsAffected > 0 {
		model := user.FindPreloadById(userModel.ID)
		r := resources.User{Model: &model}
		response.Data(c, r.ShowResource())
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *UsersController) Delete(c *gin.Context) {
	id := cast.ToUint64(c.Param("id"))
	if ok := helpers.IdVerify(id); !ok {
		response.Abort404(c)
		return
	}
	userModel := user.FindById(optimusPkg.NewOptimus().Decode(id))
	if userModel.ID == 0 {
		response.Abort404(c)
		return
	}

	rowsAffected := userModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}

func (ctrl *UsersController) Menu(c *gin.Context) {
	id := cast.ToUint64(auth.CurrentUID(c))
	if ok := helpers.IdVerify(id); !ok {
		response.Abort404(c)
		return
	}
	userModel := user.FindUserMenu(optimusPkg.NewOptimus().Decode(id))
	if userModel.ID == 0 {
		response.Abort404(c)
		return
	}
	r := resources.User{Menu: &userModel}
	data := map[string]any{
		"menus": r.MenuResource(),
	}

	response.Data(c, data)
}

func (ctrl *UsersController) UpdateAvatar(c *gin.Context) {

	request := requests.UserUpdateAvatarRequest{}
	if ok := requests.Validate(c, &request, requests.UserUpdateAvatar); !ok {
		return
	}

	id := cast.ToUint64(auth.CurrentUID(c))
	if ok := helpers.IdVerify(id); !ok {
		response.Abort404(c)
		return
	}
	userModel := user.FindById(optimusPkg.NewOptimus().Decode(id))
	if userModel.ID == 0 {
		response.Abort404(c)
		return
	}

	avatar, err := file.SaveUploadAvatar(c, request.Avatar)
	if err != nil {
		response.Abort500(c, "上传头像失败，请稍后尝试~")
		return
	}

	userModel.Avatar = config.GetString("app.url") + avatar
	rowsAffected := userModel.Save()
	if rowsAffected > 0 {
		model := user.FindById(userModel.ID)
		r := resources.User{Model: &model}
		response.Data(c, r.ShowResource())
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}
