package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"item-server/app/http/resources"
	"item-server/app/models"
	"item-server/app/models/category"
	"item-server/app/models/unit"
	"item-server/app/replaces"
	"item-server/app/requests"
	"item-server/pkg/helpers"
	optimusPkg "item-server/pkg/optimus"
	"item-server/pkg/response"
)

type UnitsController struct {
	BaseAPIController
}

func (ctrl *UnitsController) InitialValue(c *gin.Context) {
	r := resources.Category{ModelTree: category.TreeCategoryAll()}
	data := map[string]any{
		"states":     models.InitState(),
		"categories": r.CategorySelectResource(),
	}

	response.Data(c, data)
}

func (ctrl *UnitsController) Index(c *gin.Context) {

	requestFilter := requests.UnitFilterRequest{}
	if ok := requests.Validate(c, &requestFilter, requests.UnitFilter); !ok {
		return
	}
	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	//查询参数处理
	rep := replaces.UnitIndex{}
	if ok := rep.UnitIndexReplace(&requestFilter); ok != nil {
		return
	}

	data, pager := unit.Paginate(c, 10, &rep)
	r := resources.Unit{ModelSlice: data}
	response.JSON(c, gin.H{
		"data":  r.IndexResource(),
		"pager": pager,
	})
}

func (ctrl *UnitsController) Show(c *gin.Context) {
	id := cast.ToUint64(c.Param("id"))
	if ok := helpers.IdVerify(id); !ok {
		return
	}
	unitModel := unit.FindPreloadById(optimusPkg.NewOptimus().Decode(id))
	if unitModel.ID == 0 {
		response.Abort404(c)
		return
	}
	r := resources.Unit{Model: &unitModel}
	response.Data(c, r.ShowResource())
}

func (ctrl *UnitsController) Store(c *gin.Context) {

	request := requests.UnitRequest{}
	if ok := requests.Validate(c, &request, requests.UnitCreate); !ok {
		return
	}

	//分类参数处理
	rep := replaces.UnitStore{}
	if ok := rep.UnitStoreReplace(&request); ok != nil {
		return
	}
	unitModel := unit.Unit{
		State:       request.State,
		Title:       request.Title,
		Description: request.Description,
		Sort:        request.Sort,
		IsPublic:    rep.Public,
		Abbr:        rep.Abbr,
	}
	id := unitModel.CreateMany(rep.Category)
	if id > 0 {
		model := unit.FindPreloadById(id)
		r := resources.Unit{Model: &model}
		response.Created(c, r.ShowResource())
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *UnitsController) Update(c *gin.Context) {
	opt := optimusPkg.NewOptimus()
	id := cast.ToUint64(c.Param("id"))
	if ok := helpers.IdVerify(id); !ok {
		return
	}

	unitModel := unit.FindById(opt.Decode(id))
	if unitModel.ID == 0 {
		response.Abort404(c)
		return
	}

	request := requests.UnitRequest{}
	if ok := requests.Validate(c, &request, requests.UnitSave); !ok {
		return
	}
	//分类参数处理
	rep := replaces.UnitSave{}
	if ok := rep.UnitSaveReplace(&request); ok != nil {
		return
	}
	unitModel.State = request.State
	unitModel.Title = request.Title
	unitModel.Description = request.Description
	unitModel.Sort = request.Sort
	unitModel.IsPublic = rep.IsPublic
	unitModel.Abbr = rep.Abbr
	unitModel.UpdatedAt = rep.UpdatedAt
	rowsAffected := unitModel.SaveMany(rep.SelectSlice, rep.Ids)
	if rowsAffected > 0 {
		model := unit.FindPreloadById(unitModel.ID)
		r := resources.Unit{Model: &model}
		response.Data(c, r.ShowResource())
	} else {
		response.Abort500(c, "更新失败，请稍后尝试~")
	}
}

func (ctrl *UnitsController) Delete(c *gin.Context) {

	id := cast.ToUint64(c.Param("id"))
	if ok := helpers.IdVerify(id); !ok {
		return
	}

	unitModel := unit.FindById(optimusPkg.NewOptimus().Decode(id))
	if unitModel.ID == 0 {
		response.Abort404(c)
		return
	}

	rowsAffected := unitModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败，请稍后尝试~")
}
