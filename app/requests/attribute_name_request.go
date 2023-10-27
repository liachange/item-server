package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type AttributeNameRequest struct {
	State       uint8    `valid:"state" json:"state,omitempty" select:"state"`
	Genre       uint8    `valid:"genre" json:"genre,omitempty" select:"genre"`
	Category    []uint64 `valid:"category" json:"category,omitempty"`
	Title       string   `valid:"title" json:"title,omitempty" select:"title"`
	Description string   `valid:"description" json:"description,omitempty" select:"description,null"`
	Sort        uint64   `valid:"sort" json:"sort,omitempty" select:"sort"`
	Search      string   `valid:"search"  json:"search,omitempty" select:"search,null"`
}

type AttributeNameFilterRequest struct {
	State    uint8  `valid:"state" json:"state,omitempty" form:"state"`
	BetTime  string `valid:"bet_time" json:"bet_time,omitempty" form:"bet_time"`
	Genre    uint8  `valid:"genre" json:"genre,omitempty" form:"genre"`
	Title    string `valid:"title" json:"title,omitempty" form:"title"`
	Category string `valid:"category" json:"category,omitempty"  form:"category"`
	Public   uint8  `valid:"public" json:"public,omitempty" form:"public" `
}

func AttributeNameCreate(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"state":       []string{"required", "in:1,2,3"},
		"genre":       []string{"required", "in:1,2,3"},
		"title":       []string{"required", "min_cn:1", "max_cn:30", "not_exists:brands,title"},
		"description": []string{"max_cn:255"},
		"search":      []string{"max_cn:50"},
		"sort":        []string{"numeric_between:1,9999"},
		"category":    []string{"required"},
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
			"max_cn:简介长度需小于 100 个字",
		},
		"genre": []string{
			"required:状态必填项",
			"in:状态格式不正确",
		},
		"search": []string{
			"max_cn:简介长度需小于 50 个字",
		},
		"sort": []string{
			"numeric_between:排序格式不正确",
		},
		"category": []string{
			"required:分类为必填项",
		},
	}
	return validate(data, rules, messages)
}
func AttributeNameSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"state":       []string{"required", "in:1,2,3"},
		"genre":       []string{"required", "in:1,2,3"},
		"title":       []string{"required", "min_cn:1", "max_cn:30", "not_exists:brands,title," + c.Param("id")},
		"description": []string{"max_cn:255"},
		"search":      []string{"max_cn:50"},
		"sort":        []string{"numeric_between:1,9999"},
		"category":    []string{"required"},
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
			"max_cn:简介长度需小于 100 个字",
		},
		"genre": []string{
			"required:状态必填项",
			"in:状态格式不正确",
		},
		"search": []string{
			"max_cn:简介长度需小于 50 个字",
		},
		"sort": []string{
			"numeric_between:排序格式不正确",
		},
		"category": []string{
			"required:分类为必填项",
		},
	}
	return validate(data, rules, messages)
}

func AttributeNameFilter(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"state":    []string{"in:1,2,3"},
		"genre":    []string{"in:1,2,3"},
		"bet_time": []string{"slice_time"},
		"category": []string{"exist_key_str:categories"},
		"public":   []string{"in:1,2"},
	}
	messages := govalidator.MapData{
		"state": []string{
			"in:状态格式不正确",
		},
		"genre": []string{
			"in:状态格式不正确",
		},
		"bet_time": []string{
			"slice_time:时间格式不正确",
		},
		"category": []string{
			"exist_key_str:分类格式不正确",
		},
		"public": []string{
			"in:属性格式不正确",
		},
	}
	return validate(data, rules, messages)
}
