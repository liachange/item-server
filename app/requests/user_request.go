package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"item-server/app/requests/validators"
)

type UserRequest struct {
	State           uint8    `valid:"state" json:"state,omitempty" select:"state"`
	Name            string   `valid:"name" json:"name,omitempty" select:"name"`
	Password        string   `valid:"password" json:"password,omitempty" select:"password"`
	PasswordConfirm string   `valid:"password_confirm" json:"password_confirm,omitempty"`
	Role            []uint64 `valid:"role" json:"role,omitempty"`
	Phone           string   `valid:"phone" json:"phone,omitempty" select:"phone"`
	Email           string   `valid:"email" json:"email,omitempty" select:"email"`
	Nickname        string   `valid:"nickname" json:"nickname,omitempty" select:"nickname"`
}

type UserFilterRequest struct {
	State   uint8  `valid:"state" json:"state,omitempty" form:"state" filter:"state,eq"`
	Name    string `valid:"name" json:"name,omitempty" form:"name" filter:"name,like"`
	BetTime string `valid:"bet_time" json:"bet_time,omitempty" form:"bet_time" filter:"created_at,bet_time"`
}

func UserCreate(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"state":            []string{"required", "in:1,2,3"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
		"role":             []string{"required", "exist_key:roles"},
		"name":             []string{"min:3", "max:20", "exist_user:users"},
		"phone":            []string{"digits:11", "exist_user:users"},
		"email":            []string{"min:4", "max:30", "email", "exist_user:users"},
		"nickname":         []string{"min_cn:1", "max_cn:20"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"required:账户为必填项",
			"min:账户长度需至少 3 个字符",
			"max:账户长度不能超过 20 个字符",
			"exist_user:账户已存在",
		},
		"state": []string{
			"required:状态为必填项",
			"in:状态格式不正确",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"password_confirm": []string{
			"required:确认密码框为必填项",
		},
		"role": []string{
			"required:角色为必选项",
			"exist_key:角色参数异常",
		},
		"email": []string{
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
			"exist_user:Email 已存在",
		},
		"phone": []string{
			"digits:手机号长度必须为 11 位的数字",
			"exist_user:手机号已存在",
		},
		"nickname": []string{
			"min_cn:昵称 长度需大于 1",
			"max_cn:昵称 长度需小于 20",
		},
	}
	errs := validate(data, rules, messages)
	_data := data.(*UserRequest)
	errs = validators.ValidatePasswordConfirm(_data.Password, _data.PasswordConfirm, errs)

	fields := map[string]string{
		"name":  _data.Name,
		"phone": _data.Phone,
		"email": _data.Email,
	}
	errs = validators.ValidateRequiredWithout(fields, "name,phone,email其中必须一个有值", errs)
	return errs
}

func UserSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"name":             []string{"min:3", "max:20", "exist_user:users," + c.Param("id")},
		"phone":            []string{"digits:11", "exist_user:users," + c.Param("id")},
		"email":            []string{"min:4", "max:30", "email", "exist_user:users," + c.Param("id")},
		"state":            []string{"required", "in:1,2,3"},
		"password":         []string{"min:6"},
		"password_confirm": []string{"min:6"},
		"role":             []string{"required", "exist_key:roles"},
		"nickname":         []string{"min_cn:1", "max_cn:20"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"min:账户长度需至少 3 个字符",
			"max:账户长度不能超过 20 个字符",
			"exist_user:账户已存在",
		},
		"state": []string{
			"required:状态为必填项",
			"in:状态格式不正确",
		},
		"password": []string{
			"required:密码为必填项",
			"min:密码长度需大于 6",
		},
		"password_confirm": []string{
			"required:确认密码框为必填项",
		},
		"role": []string{
			"required:权限为必填项",
			"exist_key:参数异常",
		},
		"email": []string{
			"min:Email 长度需大于 4",
			"max:Email 长度需小于 30",
			"email:Email 格式不正确，请提供有效的邮箱地址",
			"exist_user:Email 已存在",
		},
		"phone": []string{
			"digits:手机号长度必须为 11 位的数字",
			"exist_user:手机号已存在",
		},
		"nickname": []string{
			"min_cn:昵称 长度需大于 1",
			"max_cn:昵称 长度需小于 20",
		},
	}
	errs := validate(data, rules, messages)
	_data := data.(*UserRequest)
	errs = validators.ValidatePasswordConfirm(_data.Password, _data.PasswordConfirm, errs)
	return errs
}

func UserFilter(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"name":     []string{"max:20"},
		"state":    []string{"in:1,2,3"},
		"bet_time": []string{"slice_time"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"max:账户长度不能超过 20 个字符",
		},
		"state": []string{
			"in:状态格式不正确",
		},
		"bet_time": []string{
			"slice_time:时间格式不正确",
		},
	}
	return validate(data, rules, messages)
}
