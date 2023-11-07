package file

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"item-server/pkg/app"
	"item-server/pkg/auth"
	"item-server/pkg/helpers"
	"mime/multipart"
	"os"
	"path/filepath"
)

// Put 将数据存入文件
func Put(data []byte, to string) error {
	err := os.WriteFile(to, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Exists 判断文件是否存在
func Exists(fileToCheck string) bool {
	if _, err := os.Stat(fileToCheck); os.IsNotExist(err) {
		return false
	}
	return true
}

func SaveUploadAvatar(c *gin.Context, file *multipart.FileHeader) (string, error) {
	var avatar string
	// 确保目录存在，不存在创建
	publicPath := "public"
	dirName := fmt.Sprintf("/uploads/avatars/%s/%s/", app.TimeNowInTimezone().Format("2006/01/02"), auth.CurrentUID(c))
	os.MkdirAll(publicPath+dirName, 0755)

	// 保存文件
	fileName := randomNameFromUploadFile(file)
	// public/uploads/avatars/2021/12/22/1/abc.png
	avatarPath := publicPath + dirName + fileName
	if err := c.SaveUploadedFile(file, avatarPath); err != nil {
		return avatar, err
	}

	return avatarPath, nil
}
func SaveUploadImage(c *gin.Context, file *multipart.FileHeader) (string, error) {
	var avatar string
	// 确保目录存在，不存在创建
	publicPath := "public"
	timePath := fmt.Sprintf("/%s/", app.TimeNowInTimezone().Format("2006/01/02"))
	dirName := fmt.Sprintf("/uploads/images%s", timePath)

	os.MkdirAll(publicPath+dirName, 0755)

	// 保存文件
	fileName := randomNameFromUploadFile(file)
	// public/uploads/avatars/2021/12/22/1/abc.png
	path := publicPath + dirName + fileName
	if err := c.SaveUploadedFile(file, path); err != nil {
		return avatar, err
	}

	return "/uploads" + timePath + fileName, nil
}
func randomNameFromUploadFile(file *multipart.FileHeader) string {
	return helpers.RandomString(16) + filepath.Ext(file.Filename)
}
