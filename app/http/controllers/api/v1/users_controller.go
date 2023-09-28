package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	v1 "item-server/app/http/Resources/api/v1"
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
		"data":  v1.IndexResource(data),
		"pager": pager,
	})
}

func (ctrl *UsersController) Show(c *gin.Context) {
	userModel := user.FirstById(cast.ToUint64(c.Param("id")))
	if userModel.ID == 0 {
		response.Abort404(c)
		return
	}
	response.Data(c, v1.ShowResource(userModel))
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
		response.Data(c, v1.ShowResource(user.FirstPreloadById(userModel.ID)))
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

//func getFilterMaps(request requests.UserFilterRequest) map[string]interface{} {
//	// eq 等于 lt小于 gt大于 neq不等于 egt大于等于 elt小于等于 bet 两个值之间 in 包含在给定的值列表中
//	maps := make(map[string]interface{})
//	if request.Name != "" {
//		maps["like:name"] = request.Name
//	}
//	if request.State > 0 {
//		maps["eq:state"] = request.State
//	}
//	if request.BetTime != "" {
//		maps["bet:created_at"] = request.BetTime
//	}
//	return maps
//}
