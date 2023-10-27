package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"item-server/app/http/resources"
	"item-server/app/models"
	"item-server/app/models/brand"
	"item-server/app/models/category"
	"item-server/app/replaces"
	"item-server/app/requests"
	"item-server/pkg/helpers"
	optimusPkg "item-server/pkg/optimus"
	"item-server/pkg/response"
)

type BrandsController struct {
	BaseAPIController
}

func (ctrl *BrandsController) InitialValue(c *gin.Context) {
	r := resources.Category{ModelTree: category.TreeCategoryAll()}
	data := map[string]any{
		"states":     models.InitState(),
		"categories": r.TreeIterative(0),
	}

	response.Data(c, data)
}

func (ctrl *BrandsController) Index(c *gin.Context) {

	requestFilter := requests.BrandFilterRequest{}
	if ok := requests.Validate(c, &requestFilter, requests.BrandFilter); !ok {
		return
	}

	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	//请求参数处理

	rep := replaces.BrandIndex{}
	if ok := rep.BrandIndexReplace(&requestFilter); ok != nil {
		return
	}

	data, pager := brand.Paginate(c, 10, &rep)
	r := resources.Brand{ModelSlice: data}
	response.JSON(c, gin.H{
		"data":  r.IndexResource(),
		"pager": pager,
	})
}

func (ctrl *BrandsController) Show(c *gin.Context) {
	id := cast.ToUint64(c.Param("id"))
	if ok := helpers.IdVerify(id); !ok {
		return
	}
	brandModel := brand.FindPreloadById(optimusPkg.NewOptimus().Decode(id))
	if brandModel.ID == 0 {
		response.Abort404(c)
		return
	}
	r := resources.Brand{Model: &brandModel}
	response.Data(c, r.ShowResource())
}

func (ctrl *BrandsController) Store(c *gin.Context) {

	request := requests.BrandRequest{}
	if ok := requests.Validate(c, &request, requests.BrandCreate); !ok {
		return
	}

	//请求参数处理
	rep := replaces.BrandStore{}
	if err := rep.BrandStoreReplace(&request); err != nil {
		return
	}

	brandModel := brand.Brand{
		State:       request.State,
		Title:       request.Title,
		Description: request.Description,
		Sort:        request.Sort,
		IconUrl:     request.IconUrl,
		IsPublic:    rep.IsPublic,
	}

	brandModel.CreateMany(rep.Ids)
	if brandModel.ID > 0 {
		model := brand.FindPreloadById(brandModel.ID)
		r := resources.Brand{Model: &model}
		response.Created(c, r.ShowResource())
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *BrandsController) Update(c *gin.Context) {

	id := cast.ToUint64(c.Param("id"))
	if ok := helpers.IdVerify(id); !ok {
		return
	}

	brandModel := brand.FindById(optimusPkg.NewOptimus().Decode(id))
	if brandModel.ID == 0 {
		response.Abort404(c)
		return
	}

	request := requests.BrandRequest{}
	if ok := requests.Validate(c, &request, requests.BrandSave); !ok {
		return
	}

	//参数处理
	rep := replaces.UpdateReplace{}
	if err := rep.BrandUpdateReplace(&request); err != nil {
		return
	}

	brandModel.State = request.State
	brandModel.Title = request.Title
	brandModel.Description = request.Description
	brandModel.IconUrl = request.IconUrl
	brandModel.IsPublic = rep.IsPublic
	brandModel.UpdatedAt = rep.UpdatedAt
	rowsAffected := brandModel.SaveMany(rep.SelectSlice, rep.Ids)
	if rowsAffected > 0 {
		model := brand.FindPreloadById(brandModel.ID)
		r := resources.Brand{Model: &model}
		response.Data(c, r.ShowResource())
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *BrandsController) Delete(c *gin.Context) {

	id := cast.ToUint64(c.Param("id"))
	if ok := helpers.IdVerify(id); !ok {
		return
	}

	brandModel := brand.FindById(optimusPkg.NewOptimus().Decode(id))
	if brandModel.ID == 0 {
		response.Abort404(c)
		return
	}

	rowsAffected := brandModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}
