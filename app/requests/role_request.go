package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type RoleRequest struct {
	State       uint8    `valid:"state" json:"state,omitempty" select:"state"`
	Name        string   `valid:"name" json:"name,omitempty" select:"name"`
	Title       string   `valid:"title" json:"title,omitempty" select:"title"`
	Description string   `valid:"description" json:"description,omitempty" select:"description"`
	Guard       string   `valid:"guard" json:"guard,omitempty" select:"guard_name"`
	Permission  []uint64 `valid:"permission" json:"permission,omitempty"`
}

type RoleFilterRequest struct {
	State   uint8  `valid:"state" json:"state,omitempty" form:"state" filter:"state,eq"`
	Name    string `valid:"name" json:"name,omitempty" form:"name" filter:"name,like"`
	Title   string `valid:"title" json:"title,omitempty" form:"title" filter:"title,like"`
	BetTime string `valid:"bet_time" json:"bet_time,omitempty" form:"bet_time" filter:"created_at,bet_time"`
}

func RoleCreate(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"state":       []string{"required", "in:1,2,3"},
		"title":       []string{"required", "min_cn:2", "max_cn:100"},
		"name":        []string{"required", "min_cn:3", "max_cn:100", "not_exists:roles,name"},
		"description": []string{"min_cn:3", "max_cn:255"},
		"guard":       []string{"required", "in:web,mobile"},
		"permission":  []string{"required"},
	}
	messages := govalidator.MapData{
		"state": []string{
			"required:状态必填项",
			"in:状态格式不正确",
		},
		"title": []string{
			"required:副名称为必填项",
			"min_cn:副名称长度需大于 2 个字",
			"max_cn:副名称长度需小于 100 个字",
		},
		"name": []string{
			"required:名称为必填项",
			"min_cn:名称长度需大于 3 个字",
			"max_cn:名称长度需小于 100 个字",
			"not_exists:名称已存在",
		},
		"guard": []string{
			"required:守卫名称为必填项",
			"in:守卫名称格式不正确",
		},
		"description": []string{
			"min_cn:简介长度需大于 3 个字",
			"max_cn:简介长度需小于 100 个字",
		},
		"permission": []string{
			"required:权限为必填项",
		},
	}
	return validate(data, rules, messages)
}
func RoleSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"state":       []string{"required", "in:1,2,3"},
		"title":       []string{"required", "min_cn:2", "max_cn:100"},
		"name":        []string{"required", "min_cn:3", "max_cn:100", "not_exists:roles,name," + c.Param("id")},
		"description": []string{"min_cn:3", "max_cn:255"},
		"guard":       []string{"required", "in:web,mobile"},
		"permission":  []string{"required"},
	}
	messages := govalidator.MapData{
		"state": []string{
			"required:状态必填项",
			"in:状态格式不正确",
		},
		"title": []string{
			"required:副名称为必填项",
			"min_cn:副名称长度需大于 2 个字",
			"max_cn:副名称长度需小于 100 个字",
		},
		"name": []string{
			"required:名称为必填项",
			"min_cn:名称长度需大于 3 个字",
			"max_cn:名称长度需小于 100 个字",
			"not_exists:名称已存在",
		},
		"guard": []string{
			"required:守卫名称为必填项",
			"in:守卫名称格式不正确",
		},
		"description": []string{
			"min_cn:简介长度需大于 3 个字",
			"max_cn:简介长度需小于 100 个字",
		},
		"permission": []string{
			"required:权限为必填项",
		},
	}
	return validate(data, rules, messages)
}

func RoleFilter(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"name":     []string{"max_cn:20"},
		"state":    []string{"in:1,2,3"},
		"bet_time": []string{"slice_time"},
		"title":    []string{"max_cn:20"},
	}
	messages := govalidator.MapData{
		"name": []string{
			"max_cn:权限地址长度不能超过 20 个字符",
		},
		"title": []string{
			"max_cn:权限名称长度不能超过 20 个字符",
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
