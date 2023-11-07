package resources

import (
	"item-server/app/models/unit"
	"item-server/pkg/helpers"
	optimusPkg "item-server/pkg/optimus"
)

type Unit struct {
	Model      *unit.Unit
	ModelSlice []unit.Unit
}

type UnitResource struct {
	ID          uint64         `json:"id"`
	State       uint8          `json:"state"`
	IsPublic    uint8          `json:"public"`
	Title       string         `json:"title"`
	Sort        uint64         `json:"sort"`
	Description string         `json:"description"`
	Categories  []*CategoryHas `json:"categories"`
	CreatedAt   string         `json:"created_at"`
	UpdatedAt   string         `json:"updated_at"`
}

func (p *Unit) ShowResource() (show UnitResource) {
	optimus := optimusPkg.NewOptimus()
	show.ID = optimus.Encode(p.Model.ID)
	show.State = p.Model.State
	show.IsPublic = p.Model.IsPublic
	show.Title = p.Model.Title
	show.Description = p.Model.Description
	show.Sort = p.Model.Sort
	show.Categories = CategoryHasResource(p.Model.Category)
	show.CreatedAt = helpers.TimeFormat(p.Model.CreatedAt, "second")
	show.UpdatedAt = helpers.TimeFormat(p.Model.UpdatedAt, "second")
	return
}
func (p *Unit) IndexResource() (index []*UnitResource) {
	optimus := optimusPkg.NewOptimus()

	for _, model := range p.ModelSlice {
		index = append(index, &UnitResource{
			ID:          optimus.Encode(model.ID),
			State:       model.State,
			IsPublic:    model.IsPublic,
			Title:       model.Title,
			Description: model.Description,
			Categories:  CategoryHasResource(model.Category),
			Sort:        model.Sort,
			CreatedAt:   helpers.TimeFormat(model.CreatedAt, "second"),
			UpdatedAt:   helpers.TimeFormat(model.UpdatedAt, "second"),
		})
	}
	return
}

type UnitSelect struct {
	ID    uint64 `json:"value"`
	Title string `json:"title"`
}

func (p *Unit) InitialResource() (sel []*UnitSelect) {
	optimus := optimusPkg.NewOptimus()

	for _, model := range p.ModelSlice {
		sel = append(sel, &UnitSelect{
			ID:    optimus.Encode(model.ID),
			Title: model.Title,
		})
	}
	return
}
