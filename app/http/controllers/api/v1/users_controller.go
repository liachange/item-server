package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"item-server/app/http/resources"
	"item-server/app/models/user"
	"item-server/app/requests"
	"item-server/pkg/hash"
	"item-server/pkg/helpers"
	"item-server/pkg/response"
)

type UsersController struct {
	BaseAPIController
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
	response.JSON(c, gin.H{
		"data":  resources.UserIndexResource(data),
		"pager": pager,
	})
}

func (ctrl *UsersController) Show(c *gin.Context) {
	userModel := user.FirstById(cast.ToUint64(c.Param("id")))
	if userModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, resources.UserShowResource(userModel))
}

func (ctrl *UsersController) Store(c *gin.Context) {

	request := requests.UserRequest{}
	if ok := requests.Validate(c, &request, requests.UserCreate); !ok {
		return
	}

	userModel := user.User{
		Name:     request.Name,
		Password: request.Password,
		State:    request.State,
		Nickname: request.Nickname,
		Phone:    request.Phone,
		Email:    request.Email,
	}
	userModel.CreateMany(request.Role)
	if userModel.ID > 0 {
		response.Created(c, userModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *UsersController) Update(c *gin.Context) {

	userModel := user.FirstById(cast.ToUint64(c.Param("id")))
	if userModel.ID == 0 {
		response.Abort404(c)
		return
	}
	request := requests.UserRequest{}
	if ok := requests.Validate(c, &request, requests.UserSave); !ok {
		return
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
	rowsAffected := userModel.SaveMany(&request, request.Role)
	if rowsAffected > 0 {
		response.Data(c, resources.UserShowResource(user.FirstPreloadById(userModel.ID)))
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *UsersController) Delete(c *gin.Context) {

	userModel := user.FirstById(cast.ToUint64(c.Param("id")))
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
