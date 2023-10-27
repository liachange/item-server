package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"mime/multipart"
)

type ImageRequest struct {
	Image *multipart.FileHeader `valid:"image" form:"image"`
}

func ImageUpload(data interface{}, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		// size 的单位为 bytes
		// - 1024 bytes 为 1kb
		// - 1048576 bytes 为 1mb
		// - 20971520 bytes 为 20mb
		"file:image": []string{"ext:png,jpg,jpeg", "size:1048576", "required"},
	}
	messages := govalidator.MapData{
		"file:image": []string{
			"ext:只能上传 png, jpg, jpeg 任意一种的图片",
			"size:头像文件最大不能超过 20MB",
			"required:必须上传图片",
		},
	}

	return validateFile(c, data, rules, messages)
}
