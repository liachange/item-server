package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"item-server/app/http/resources"
	"item-server/app/models"
	attributeName "item-server/app/models/attribute_name"
	"item-server/app/models/category"
	"item-server/app/replaces"
	"item-server/app/requests"
	"item-server/pkg/helpers"
	optimusPkg "item-server/pkg/optimus"
	"item-server/pkg/response"
)

type AttributeNamesController struct {
	BaseAPIController
}

func (ctrl *AttributeNamesController) InitialValue(c *gin.Context) {
	r := resources.Category{ModelTree: category.TreeCategoryAll()}
	data := map[string]any{
		"states":     models.InitState(),
		"genres":     attributeName.InitGenre(),
		"categories": r.CategorySelectResource(),
	}

	response.Data(c, data)
}
func (ctrl *AttributeNamesController) Index(c *gin.Context) {

	requestFilter := requests.AttributeNameFilterRequest{}
	if ok := requests.Validate(c, &requestFilter, requests.AttributeNameFilter); !ok {
		return
	}

	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	//查询参数处理
	rep := replaces.AttributeNameIndex{}
	if ok := rep.AttributeNameReplace(&requestFilter); ok != nil {
		return
	}

	data, pager := attributeName.Paginate(c, 10, &rep)
	r := resources.AttributeName{ModelSlice: data}
	response.JSON(c, gin.H{
		"data":  r.IndexResource(),
		"pager": pager,
	})
}

func (ctrl *AttributeNamesController) Show(c *gin.Context) {
	id := cast.ToUint64(c.Param("id"))
	if ok := helpers.IdVerify(id); !ok {
		return
	}
	attributeNameModel := attributeName.FindPreloadById(optimusPkg.NewOptimus().Decode(id))
	if attributeNameModel.ID == 0 {
		response.Abort404(c)
		return
	}
	r := resources.AttributeName{Model: &attributeNameModel}
	response.Data(c, r.ShowResource())
}

func (ctrl *AttributeNamesController) Store(c *gin.Context) {

	request := requests.AttributeNameRequest{}
	if ok := requests.Validate(c, &request, requests.AttributeNameCreate); !ok {
		return
	}

	rep := replaces.AttributeNameStore{}
	if err := rep.AttributeNameStoreReplace(&request); err != nil {
		return
	}

	attributeNameModel := attributeName.AttributeName{
		State:       request.State,
		Title:       request.Title,
		Description: request.Description,
		Sort:        request.Sort,
		Genre:       request.Genre,
		IsPublic:    rep.Public,
		Abbr:        rep.Abbr,
		Search:      request.Search,
	}
	attributeNameModel.CreateMany(rep.Category)

	if attributeNameModel.ID > 0 {
		model := attributeName.FindPreloadById(attributeNameModel.ID)
		r := resources.AttributeName{Model: &model}
		response.Created(c, r.ShowResource())
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *AttributeNamesController) Update(c *gin.Context) {

	id := cast.ToUint64(c.Param("id"))
	if ok := helpers.IdVerify(id); !ok {
		return
	}
	attributeNameModel := attributeName.FindById(optimusPkg.NewOptimus().Decode(id))
	if attributeNameModel.ID == 0 {
		response.Abort404(c)
		return
	}

	request := requests.AttributeNameRequest{}
	if ok := requests.Validate(c, &request, requests.AttributeNameSave); !ok {
		return
	}

	//参数处理
	rep := replaces.AttributeNameSave{}
	if ok := rep.AttributeNameSaveReplace(&request); ok != nil {
		return
	}
	attributeNameModel.State = request.State
	attributeNameModel.Genre = request.Genre
	attributeNameModel.Title = request.Title
	attributeNameModel.IsPublic = rep.IsPublic
	attributeNameModel.Abbr = rep.Abbr
	attributeNameModel.UpdatedAt = rep.UpdatedAt
	attributeNameModel.Description = request.Description
	attributeNameModel.Search = request.Search
	rowsAffected := attributeNameModel.SaveMany(rep.SelectSlice, rep.Ids)

	if rowsAffected > 0 {
		model := attributeName.FindPreloadById(attributeNameModel.ID)
		r := resources.AttributeName{Model: &model}
		response.Data(c, r.ShowResource())
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *AttributeNamesController) Delete(c *gin.Context) {

	id := cast.ToUint64(c.Param("id"))
	if ok := helpers.IdVerify(id); !ok {
		return
	}

	attributeNameModel := attributeName.FindById(optimusPkg.NewOptimus().Decode(id))
	if attributeNameModel.ID == 0 {
		response.Abort404(c)
		return
	}

	rowsAffected := attributeNameModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}
