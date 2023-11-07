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
		"categories": r.CategorySelectResource(),
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

	//请求参数处理
	rep := replaces.CategoryStore{}
	if err := rep.CategoryStoreReplace(&request); err != nil {
		return
	}

	categoryModel := category.Category{
		State:       request.State,
		Title:       request.Title,
		Description: request.Description,
		ParentId:    rep.ParentId,
		Sort:        request.Sort,
		IconUrl:     request.IconUrl,
		Abbr:        rep.Abbr,
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

	//请求参数处理
	rep := replaces.CategorySave{}
	if err := rep.CategorySaveReplace(&request); err != nil {
		return
	}

	categoryModel.State = request.State
	categoryModel.Title = request.Title
	categoryModel.Description = request.Description
	categoryModel.IconUrl = request.IconUrl
	categoryModel.ParentId = rep.ParentId
	categoryModel.Abbr = rep.Abbr
	categoryModel.UpdatedAt = helpers.TimeNow()
	rowsAffected := categoryModel.Save(rep.SelectSlice)
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
