package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type UnitRequest struct {
	State       uint8    `valid:"state" json:"state,omitempty" select:"state"`
	Title       string   `valid:"title" json:"title,omitempty" select:"title"`
	Description string   `valid:"description" json:"description,omitempty" select:"description,null"`
	Sort        uint64   `valid:"sort" json:"sort,omitempty" select:"sort"`
	Category    []uint64 `valid:"category" json:"category,omitempty"`
}

type UnitFilterRequest struct {
	State    uint8  `valid:"state" json:"state,omitempty" form:"state"`
	Title    string `valid:"title" json:"title,omitempty" form:"title"`
	Category string `valid:"category" json:"category,omitempty"  form:"category"`
	Public   uint8  `valid:"public" json:"public,omitempty" form:"public" `
	BetTime  string `valid:"bet_time" json:"bet_time,omitempty" form:"bet_time"`
}

func UnitCreate(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"title":       []string{"required", "min_cn:1", "max_cn:30", "not_exists:categories,title"},
		"description": []string{"max_cn:255"},
		"sort":        []string{"numeric_between:1,9999"},
		"category":    []string{"required", "exist_key:格式不正确"},
	}
	messages := govalidator.MapData{
		"state": []string{
			"required:状态必填项",
			"in:状态格式不正确",
		},
		"title": []string{
			"required:名称为必填项",
			"min_cn:名称长度需大于 1 个字",
			"max_cn:名称长度需小于 30 个字",
			"not_exists:名称已存在",
		},
		"sort": []string{
			"numeric_between:排序格式不正确",
		},
		"category": []string{
			"required:分类必填项",
			"exist_key:格式不正确",
		},
	}
	return validate(data, rules, messages)
}
func UnitSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"title":       []string{"required", "min_cn:1", "max_cn:30", "not_exists:categories,title," + c.Param("id")},
		"description": []string{"max_cn:255"},
		"sort":        []string{"numeric_between:1,9999"},
		"category":    []string{"required", "exist_key:格式不正确"},
	}
	messages := govalidator.MapData{
		"state": []string{
			"required:状态必填项",
			"in:状态格式不正确",
		},
		"title": []string{
			"required:名称为必填项",
			"min_cn:名称长度需大于 1 个字",
			"max_cn:名称长度需小于 30 个字",
			"not_exists:名称已存在",
		},
		"sort": []string{
			"numeric_between:排序格式不正确",
		},
		"category": []string{
			"required:分类必填项",
			"exist_key:格式不正确",
		},
	}
	return validate(data, rules, messages)
}

func UnitFilter(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"state":    []string{"in:1,2,3"},
		"bet_time": []string{"slice_time"},
	}
	messages := govalidator.MapData{
		"state": []string{
			"in:状态格式不正确",
		},
		"bet_time": []string{
			"slice_time:时间格式不正确",
		},
	}
	return validate(data, rules, messages)
}
