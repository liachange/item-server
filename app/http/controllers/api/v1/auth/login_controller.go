package auth

import (
	"github.com/gin-gonic/gin"
	v1 "item-server/app/http/controllers/api/v1"
	"item-server/app/requests"
	"item-server/pkg/auth"
	"item-server/pkg/jwt"
	"item-server/pkg/response"
)

// LoginController 用户控制器
type LoginController struct {
	v1.BaseAPIController
}

// LoginByPhone 手机登录
func (lc *LoginController) LoginByPhone(c *gin.Context) {

	// 1. 验证表单
	request := requests.LoginByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPhone); !ok {
		return
	}

	// 2. 尝试登录
	user, err := auth.LoginByPhone(request.Phone)
	if err != nil {
		// 失败，显示错误提示
		response.Error(c, err, "账号不存在")
	} else {
		// 登录成功
		token := jwt.NewJWT().IssueToken(jwt.UserInfo{
			Nickname: user.Nickname,
			UserID:   user.GetStringID(),
		})

		response.JSON(c, gin.H{
			"token": token,
		})
	}
}
