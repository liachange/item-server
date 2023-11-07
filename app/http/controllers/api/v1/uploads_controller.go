package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"item-server/app/models/user"
	"item-server/app/requests"
	"item-server/pkg/auth"
	"item-server/pkg/file"
	"item-server/pkg/helpers"
	optimusPkg "item-server/pkg/optimus"
	"item-server/pkg/response"
)

type UploadsController struct {
	BaseAPIController
}

func (ctrl *UploadsController) Store(c *gin.Context) {

	request := requests.ImageRequest{}
	if ok := requests.Validate(c, &request, requests.ImageUpload); !ok {
		return
	}

	id := cast.ToUint64(auth.CurrentUID(c))
	if ok := helpers.IdVerify(id); !ok {
		response.Abort404(c)
		return
	}
	userModel := user.FindById(optimusPkg.NewOptimus().Decode(id))
	if userModel.ID == 0 {
		response.Abort404(c)
		return
	}

	image, err := file.SaveUploadImage(c, request.Image)
	if err != nil {
		response.Abort500(c, "文件上传失败，请稍后尝试~")
		return
	}
	imgUrl := map[string]string{
		"url": image,
	}
	response.Data(c, imgUrl)
}
func (ctrl *UploadsController) ImagePath() string {
	return "public/uploads/images"
}
