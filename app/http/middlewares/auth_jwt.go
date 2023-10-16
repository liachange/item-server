package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"item-server/app/models/user"
	"item-server/pkg/config"
	"item-server/pkg/jwt"
	optimusPkg "item-server/pkg/optimus"
	"item-server/pkg/response"
)

func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 从标头 Authorization:Bearer xxxxx 中获取信息，并验证 JWT 的准确性
		claims, err := jwt.NewJWT().ParserToken(c)

		// JWT 解析失败，有错误发生
		if err != nil {
			response.Unauthorized(c, fmt.Sprintf("请查看 %v 相关的接口认证文档", config.GetString("app.name")))
			return
		}

		// JWT 解析成功，设置用户信息
		id := cast.ToUint64(claims.UserID)
		if id > 0 {
			id = optimusPkg.NewOptimus().Decode(id)
		} else {
			response.Unauthorized(c, "找不到对应用户，用户可能已删除")
			return
		}
		userModel := user.GetById(id)
		if userModel.ID == 0 {
			response.Unauthorized(c, "找不到对应用户，用户可能已删除")
			return
		}

		// 将用户信息存入 gin.context 里，后续 auth 包将从这里拿到当前用户数据
		c.Set("current_user_id", cast.ToString(optimusPkg.NewOptimus().Encode(userModel.ID)))
		c.Set("current_user_name", userModel.Name)
		c.Set("current_user", userModel)

		c.Next()
	}
}
