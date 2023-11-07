package resources

import (
	"item-server/app/models/brand"
	"item-server/pkg/helpers"
	optimusPkg "item-server/pkg/optimus"
)

type Brand struct {
	Model      *brand.Brand
	ModelSlice []brand.Brand
}

type BrandResource struct {
	ID          uint64         `json:"id"`
	State       uint8          `json:"state"`
	Title       string         `json:"title"`
	Icon        string         `json:"icon"`
	Sort        uint64         `json:"sort"`
	Description string         `json:"description"`
	Public      uint8          `json:"public"`
	Categories  []*CategoryHas `json:"categories"`
	CreatedAt   string         `json:"created_at"`
	UpdatedAt   string         `json:"updated_at"`
}

func (p *Brand) ShowResource() (show BrandResource) {
	optimus := optimusPkg.NewOptimus()
	show.ID = optimus.Encode(p.Model.ID)
	show.Title = p.Model.Title
	show.Sort = p.Model.Sort
	show.Icon = p.Model.IconUrl
	show.State = p.Model.State
	show.Description = p.Model.Description
	show.Categories = CategoryHasResource(p.Model.Category)
	show.UpdatedAt = helpers.TimeFormat(p.Model.UpdatedAt, "second")
	show.CreatedAt = helpers.TimeFormat(p.Model.CreatedAt, "second")
	return
}

func (p *Brand) IndexResource() (index []*BrandResource) {
	optimus := optimusPkg.NewOptimus()
	for _, model := range p.ModelSlice {
		index = append(index, &BrandResource{
			ID:          optimus.Encode(model.ID),
			State:       model.State,
			Title:       model.Title,
			Description: model.Description,
			Icon:        model.IconUrl,
			Sort:        model.Sort,
			Public:      model.IsPublic,
			Categories:  CategoryHasResource(model.Category),
			CreatedAt:   helpers.TimeFormat(model.CreatedAt, "second"),
			UpdatedAt:   helpers.TimeFormat(model.UpdatedAt, "second"),
		})
	}
	return
}

type BrandSelect struct {
	ID    uint64 `json:"value"`
	Title string `json:"title"`
}

func (p *Brand) InitialResource() (sel []*BrandSelect) {
	optimus := optimusPkg.NewOptimus()
	for _, model := range p.ModelSlice {
		sel = append(sel, &BrandSelect{
			ID:    optimus.Encode(model.ID),
			Title: model.Title,
		})
	}
	return
}
