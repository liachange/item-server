package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type CategoryRequest struct {
	State       uint8  `valid:"state" json:"state,omitempty" select:"state"`
	Title       string `valid:"title" json:"title,omitempty" select:"title"`
	Description string `valid:"description" json:"description,omitempty" select:"description,null"`
	IconUrl     string `valid:"icon" json:"icon,omitempty" select:"icon_url,null"`
	Sort        uint64 `valid:"sort" json:"sort,omitempty" select:"sort"`
	ParentId    uint64 `valid:"parent" json:"parent,omitempty" select:"parent_id"`
}

type CategoryFilterRequest struct {
	Parent   uint64 `valid:"parent" json:"parent,omitempty" form:"parent"`
	State    uint8  `valid:"state" json:"state,omitempty" form:"state"`
	Title    string `valid:"title" json:"title,omitempty" form:"title"`
	BetTime  string `valid:"bet_time" json:"bet_time,omitempty" form:"bet_time"`
	Category string `valid:"category" json:"category,omitempty" form:"category"`
}

func CategoryCreate(data interface{}, c *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"state":       []string{"required", "in:1,2,3"},
		"title":       []string{"required", "min_cn:1", "max_cn:30", "not_exists:categories,title"},
		"description": []string{"max_cn:255"},
		"icon":        []string{"image_suffix", "max:255"},
		"sort":        []string{"numeric_between:1,9999"},
		"parent":      []string{"numeric", "exists:categories,id"},
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
		"description": []string{
			"min_cn:简介长度需大于 3 个字",
			"max_cn:简介长度需小于 100 个字",
		},
		"icon": []string{
			"image_suffix:图标文件格式不正确",
			"max:图标长度需小于 255 个字",
		},
		"sort": []string{
			"numeric_between:排序格式不正确",
		},
		"parent": []string{
			"numeric:父标识格式不正确",
			"exists:父标识不正确",
		},
	}
	return validate(data, rules, messages)
}
func CategorySave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"state":       []string{"required", "in:1,2,3"},
		"title":       []string{"required", "min_cn:2", "max_cn:100", "not_exists:categories,title," + c.Param("id")},
		"description": []string{"max_cn:255"},
		"icon":        []string{"image_suffix", "max:255"},
		"sort":        []string{"numeric_between:1,9999"},
		"parent":      []string{"numeric", "exists:categories,id"},
	}
	messages := govalidator.MapData{
		"state": []string{
			"required:状态必填项",
			"in:状态格式不正确",
		},
		"title": []string{
			"required:名称为必填项",
			"min_cn:名称长度需大于 2 个字",
			"max_cn:名称长度需小于 100 个字",
			"not_exists:名称已存在",
		},
		"description": []string{
			"max_cn:简介长度需小于 100 个字",
		},
		"icon": []string{
			"image_suffix:图标文件格式不正确",
			"max:图标长度需小于 255 个字",
		},
		"sort": []string{
			"numeric_between:排序格式不正确",
		},
		"parent": []string{
			"numeric:父标识格式不正确",
			"exists:父标识不正确",
		},
	}
	return validate(data, rules, messages)
}

func CategoryFilter(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"state":    []string{"in:1,2,3"},
		"bet_time": []string{"slice_time"},
		"title":    []string{"max_cn:30"},
		"category": []string{"exist_key_str:categories"},
	}
	messages := govalidator.MapData{
		"state": []string{
			"in:状态格式不正确",
		},
		"bet_time": []string{
			"slice_time:时间格式不正确",
		},
		"title": []string{
			"max_cn:权限名称长度不能超过 30 个字符",
		},
		"category": []string{
			"exist_key_str:分类格式不正确",
		},
	}
	return validate(data, rules, messages)
}
