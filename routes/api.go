package routes

import (
	"github.com/gin-gonic/gin"
	controllers "item-server/app/http/controllers/api/v1"
	"item-server/app/http/controllers/api/v1/auth"
	"item-server/app/http/middlewares"
	"item-server/pkg/config"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存放到这里
	var v1 *gin.RouterGroup
	if len(config.Get("app.api_domain")) == 0 {
		v1 = r.Group("/api/v1")
	} else {
		v1 = r.Group("/v1")
	}
	{
		authGroup := v1.Group("/auth")
		// 限流中间件：每小时限流，作为参考 Github API 每小时最多 60 个请求（根据 IP）
		// 测试时，可以调高一点
		authGroup.Use(middlewares.LimitIP("1000-H"))
		{
			suc := new(auth.SignupController)
			// 判断手机是否已注册
			authGroup.POST("/signup/phone/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.IsPhoneExist)
			// 判断 Email 是否已注册
			authGroup.POST("/signup/email/exist", middlewares.GuestJWT(), middlewares.LimitPerRoute("60-H"), suc.IsEmailExist)

			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			// 图片验证码，需要加限流
			authGroup.POST("/verify-codes/captcha", middlewares.LimitPerRoute("20-H"), vcc.ShowCaptcha)
			// 手机短信验证码
			authGroup.POST("/verify-codes/phone", middlewares.LimitPerRoute("20-H"), vcc.SendUsingPhone)
			// 邮箱验证码
			authGroup.POST("/verify-codes/email", middlewares.LimitPerRoute("50-H"), vcc.SendUsingEmail)
			//手机号注册
			authGroup.POST("/signup/using-phone", middlewares.GuestJWT(), suc.SignupUsingPhone)
			//邮箱注册
			authGroup.POST("/signup/using-email", middlewares.GuestJWT(), suc.SignupUsingEmail)

			lgc := new(auth.LoginController)
			// 使用手机号，短信验证码进行登录
			authGroup.POST("/login/using-phone", middlewares.GuestJWT(), lgc.LoginByPhone)
			// 支持手机号，Email 和 用户名
			authGroup.POST("/login/using-password", middlewares.GuestJWT(), lgc.LoginByPassword)
			//刷新令牌
			authGroup.POST("/login/refresh-token", middlewares.AuthJWT(), lgc.RefreshToken)

			// 重置密码
			pwc := new(auth.PasswordController)
			//手机号找回密码
			authGroup.POST("/password-reset/using-phone", middlewares.GuestJWT(), pwc.ResetByPhone)
			//邮箱找回密码
			authGroup.POST("/password-reset/using-email", middlewares.GuestJWT(), pwc.ResetByEmail)

			//权限
			pmc := new(controllers.PermissionsController)
			pmcGroup := v1.Group("/permissions")
			{
				pmcGroup.GET("/initial", middlewares.AuthJWT(), pmc.InitialValue)
				pmcGroup.GET("", middlewares.AuthJWT(), pmc.Index)
				pmcGroup.POST("", middlewares.AuthJWT(), pmc.Store)
				pmcGroup.PUT("/:id", middlewares.AuthJWT(), pmc.Update)
				pmcGroup.DELETE("/:id", middlewares.AuthJWT(), pmc.Delete)
				pmcGroup.GET("/:id", middlewares.AuthJWT(), pmc.Show)
			}

			//角色
			rc := new(controllers.RolesController)
			rcGroup := v1.Group("/roles")
			{
				rcGroup.GET("/initial", middlewares.AuthJWT(), rc.InitialValue)
				rcGroup.GET("", middlewares.AuthJWT(), rc.Index)
				rcGroup.POST("", middlewares.AuthJWT(), rc.Store)
				rcGroup.PUT("/:id", middlewares.AuthJWT(), rc.Update)
				rcGroup.DELETE("/:id", middlewares.AuthJWT(), rc.Delete)
				rcGroup.GET("/:id", middlewares.AuthJWT(), rc.Show)
			}
			//账户
			uc := new(controllers.UsersController)
			ucGroup := v1.Group("/users")
			{
				ucGroup.GET("/menus", middlewares.AuthJWT(), uc.Menu)
				ucGroup.GET("/initial", middlewares.AuthJWT(), uc.InitialValue)
				ucGroup.GET("", middlewares.AuthJWT(), uc.Index)
				ucGroup.POST("", middlewares.AuthJWT(), uc.Store)
				ucGroup.PUT("/:id", middlewares.AuthJWT(), uc.Update)
				ucGroup.DELETE("/:id", middlewares.AuthJWT(), uc.Delete)
				ucGroup.GET("/:id", middlewares.AuthJWT(), uc.Show)
			}
		}
	}
}
