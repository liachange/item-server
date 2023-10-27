package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"item-server/app/http/resources"
	"item-server/app/models"
	"item-server/app/models/category"
	"item-server/app/replaces"
	"item-server/app/requests"
	"item-server/pkg/helpers"
	optimusPkg "item-server/pkg/optimus"
	"item-server/pkg/response"
)

type CategoriesController struct {
	BaseAPIController
}

func (ctrl *CategoriesController) InitialValue(c *gin.Context) {
	r := resources.Category{ModelTree: category.TreeCategoryAll()}
	data := map[string]any{
		"states":     models.InitState(),
		"categories": r.TreeIterative(0),
	}

	response.Data(c, data)
}

func (ctrl *CategoriesController) Index(c *gin.Context) {

	requestFilter := requests.CategoryFilterRequest{}
	if ok := requests.Validate(c, &requestFilter, requests.CategoryFilter); !ok {
		return
	}
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	//查询参数处理
	rep := replaces.CategoryIndex{}
	if ok := rep.BrandIndexReplace(&requestFilter); ok != nil {
		return
	}

	data, pager := category.Paginate(c, 10, &rep)
	r := resources.Category{ModelSlice: data}
	response.JSON(c, gin.H{
		"data":  r.IndexResource(),
		"pager": pager,
	})
}

func (ctrl *CategoriesController) Show(c *gin.Context) {
	id := cast.ToUint64(c.Param("id"))
	if ok := helpers.IdVerify(id); !ok {
		return
	}
	categoryModel := category.FindById(optimusPkg.NewOptimus().Decode(id))
	if categoryModel.ID == 0 {
		response.Abort404(c)
		return
	}
	r := resources.Category{Model: &categoryModel}
	response.Data(c, r.ShowResource())
}

func (ctrl *CategoriesController) Store(c *gin.Context) {

	request := requests.CategoryRequest{}
	if ok := requests.Validate(c, &request, requests.CategoryCreate); !ok {
		return
	}

	var parent uint64
	if request.ParentId > 0 {
		parent = optimusPkg.NewOptimus().Decode(request.ParentId)
	}
	categoryModel := category.Category{
		State:       request.State,
		Title:       request.Title,
		Description: request.Description,
		ParentId:    parent,
		Sort:        request.Sort,
		IconUrl:     request.IconUrl,
	}
	id := categoryModel.Create()
	if id > 0 {
		model := category.FindById(id)
		r := resources.Category{Model: &model}
		response.Created(c, r.ShowResource())
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *CategoriesController) Update(c *gin.Context) {

	id := cast.ToUint64(c.Param("id"))
	if ok := helpers.IdVerify(id); !ok {
		return
	}

	categoryModel := category.FindById(optimusPkg.NewOptimus().Decode(id))
	if categoryModel.ID == 0 {
		response.Abort404(c)
		return
	}

	request := requests.CategoryRequest{}
	if ok := requests.Validate(c, &request, requests.CategorySave); !ok {
		return
	}
	parent := request.ParentId
	if parent > 0 {
		parent = optimusPkg.NewOptimus().Decode(request.ParentId)
	} else {
		parent = 0
	}
	categoryModel.State = request.State
	categoryModel.Title = request.Title
	categoryModel.Description = request.Description
	categoryModel.IconUrl = request.IconUrl
	categoryModel.ParentId = parent
	categoryModel.UpdatedAt = helpers.TimeNow()
	rowsAffected := categoryModel.Save(&request)
	if rowsAffected > 0 {
		model := category.FindById(categoryModel.ID)
		r := resources.Category{Model: &model}
		response.Data(c, r.ShowResource())
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *CategoriesController) Delete(c *gin.Context) {

	id := cast.ToUint64(c.Param("id"))
	if ok := helpers.IdVerify(id); !ok {
		return
	}

	categoryModel := category.FindById(optimusPkg.NewOptimus().Decode(id))
	if categoryModel.ID == 0 {
		response.Abort404(c)
		return
	}

	rowsAffected := categoryModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}
