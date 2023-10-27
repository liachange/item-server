package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type AttributeValueRequest struct {
	State           uint8  `valid:"state" json:"state,omitempty" select:"state"`
	AttributeNameId uint64 `valid:"attribute_name" json:"attribute_name,omitempty" select:"attribute_name_id"`
	Title           string `valid:"title" json:"title,omitempty" select:"title"`
	Description     string `valid:"description" json:"description,omitempty" select:"description,null"`
	IconUrl         string `valid:"icon" json:"icon,omitempty" select:"icon_url,null"`
	Sort            uint64 `valid:"sort" json:"sort,omitempty" select:"sort"`
	Search          string `valid:"search"  json:"search,omitempty" select:"search,null"`
}

type AttributeValueFilterRequest struct {
	State         uint8  `valid:"state" json:"state,omitempty" form:"state"`
	Title         string `valid:"title" json:"title,omitempty" form:"title"`
	AttributeName string `valid:"attribute_name" json:"attribute_name,omitempty"  form:"attribute_name"`
}

func AttributeValueCreate(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"state":          []string{"required", "in:1,2,3"},
		"title":          []string{"required", "min_cn:1", "max_cn:30", "not_exists:attribute_values,title"},
		"description":    []string{"max_cn:255"},
		"search":         []string{"max_cn:50"},
		"icon":           []string{"image_suffix", "max:255"},
		"sort":           []string{"numeric_between:1,9999"},
		"attribute_name": []string{"required", "numeric", "exists:attribute_names,id"},
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
		"icon": []string{
			"image_suffix:图标文件格式不正确",
			"max:图标长度需小于 255 个字",
		},
		"sort": []string{
			"numeric_between:排序格式不正确",
		},
		"attribute_name": []string{
			"numeric:父标识格式不正确",
			"exists:父标识不正确",
		},
	}
	return validate(data, rules, messages)
}
func AttributeValueSave(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"state":          []string{"required", "in:1,2,3"},
		"title":          []string{"required", "min_cn:1", "max_cn:30", "not_exists:attribute_values,title," + c.Param("id")},
		"description":    []string{"max_cn:255"},
		"search":         []string{"max_cn:50"},
		"icon":           []string{"image_suffix", "max:255"},
		"sort":           []string{"numeric_between:1,9999"},
		"attribute_name": []string{"numeric", "exists:attribute_names,id"},
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
		"icon": []string{
			"image_suffix:图标文件格式不正确",
			"max:图标长度需小于 255 个字",
		},
		"sort": []string{
			"numeric_between:排序格式不正确",
		},
		"attribute_name": []string{
			"numeric:父标识格式不正确",
			"exists:父标识不正确",
		},
	}
	return validate(data, rules, messages)
}

func AttributeValueFilter(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"state":          []string{"in:1,2,3"},
		"attribute_name": []string{"exist_key_str:attribute_names"},
	}
	messages := govalidator.MapData{
		"state": []string{
			"in:状态格式不正确",
		},
		"attribute_name": []string{
			"exist_key_str:属性名称格式不正确",
		},
	}
	return validate(data, rules, messages)
}
