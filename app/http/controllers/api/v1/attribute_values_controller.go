package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"item-server/app/http/resources"
	"item-server/app/models"
	attributeName "item-server/app/models/attribute_name"
	attributeValue "item-server/app/models/attribute_value"
	"item-server/app/replaces"
	"item-server/app/requests"
	"item-server/pkg/helpers"
	optimusPkg "item-server/pkg/optimus"
	"item-server/pkg/pinyin"
	"item-server/pkg/response"
)

type AttributeValuesController struct {
	BaseAPIController
}

func (ctrl *AttributeValuesController) InitialValue(c *gin.Context) {
	r := resources.AttributeName{ModelSlice: attributeName.AttributeNameAll()}
	data := map[string]any{
		"states":          models.InitState(),
		"attribute_names": r.InitialResource(),
	}

	response.Data(c, data)
}

func (ctrl *AttributeValuesController) Index(c *gin.Context) {

	requestFilter := requests.AttributeValueFilterRequest{}
	if ok := requests.Validate(c, &requestFilter, requests.AttributeValueFilter); !ok {
		return
	}

	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	rep := replaces.AttributeValueIndex{}
	if ok := rep.AttrValueReplace(&requestFilter); ok != nil {
		return
	}

	data, pager := attributeValue.Paginate(c, 10, &rep)
	r := resources.AttributeValue{ModelSlice: data}
	response.JSON(c, gin.H{
		"data":  r.IndexResource(),
		"pager": pager,
	})
}

func (ctrl *AttributeValuesController) Show(c *gin.Context) {
	id := cast.ToUint64(c.Param("id"))
	if ok := helpers.IdVerify(id); !ok {
		return
	}
	attributeValueModel := attributeValue.FindById(optimusPkg.NewOptimus().Decode(id))
	if attributeValueModel.ID == 0 {
		response.Abort404(c)
		return
	}
	r := resources.AttributeValue{Model: &attributeValueModel}
	response.Data(c, r.ShowResource())
}

func (ctrl *AttributeValuesController) Store(c *gin.Context) {

	request := requests.AttributeValueRequest{}
	if ok := requests.Validate(c, &request, requests.AttributeValueCreate); !ok {
		return
	}

	attributeValueModel := attributeValue.AttributeValue{
		State:           request.State,
		Title:           request.Title,
		Description:     request.Description,
		AttributeNameId: optimusPkg.NewOptimus().Decode(request.AttributeNameId),
		Sort:            request.Sort,
		Abbr:            pinyin.GetFirstSpell(request.Title),
		Search:          request.Search,
	}
	attributeValueModel.Create()
	if attributeValueModel.ID > 0 {
		model := attributeValue.FindById(attributeValueModel.ID)
		r := resources.AttributeValue{Model: &model}
		response.Created(c, r.ShowResource())
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *AttributeValuesController) Update(c *gin.Context) {

	id := cast.ToUint64(c.Param("id"))
	if ok := helpers.IdVerify(id); !ok {
		return
	}

	attributeValueModel := attributeValue.FindById(optimusPkg.NewOptimus().Decode(id))
	if attributeValueModel.ID == 0 {
		response.Abort404(c)
		return
	}

	request := requests.AttributeValueRequest{}
	if ok := requests.Validate(c, &request, requests.AttributeValueSave); !ok {
		return
	}

	attributeValueModel.State = request.State
	attributeValueModel.Title = request.Title
	attributeValueModel.Description = request.Description
	attributeValueModel.AttributeNameId = optimusPkg.NewOptimus().Decode(request.AttributeNameId)
	attributeValueModel.Search = request.Search
	attributeValueModel.Abbr = pinyin.GetFirstSpell(request.Title)
	attributeValueModel.UpdatedAt = helpers.TimeNow()
	rowsAffected := attributeValueModel.Save(&request)
	if rowsAffected > 0 {
		model := attributeValue.FindById(attributeValueModel.ID)
		r := resources.AttributeValue{Model: &model}
		response.Data(c, r.ShowResource())
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *AttributeValuesController) Delete(c *gin.Context) {

	id := cast.ToUint64(c.Param("id"))
	if ok := helpers.IdVerify(id); !ok {
		return
	}

	attributeValueModel := attributeValue.FindById(optimusPkg.NewOptimus().Decode(id))
	if attributeValueModel.ID == 0 {
		response.Abort404(c)
		return
	}

	rowsAffected := attributeValueModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}
