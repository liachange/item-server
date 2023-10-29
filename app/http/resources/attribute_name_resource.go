package resources

import (
	"item-server/app/models/attribute_name"
	"item-server/pkg/helpers"
	optimusPkg "item-server/pkg/optimus"
)

type AttributeName struct {
	Model      *attribute_name.AttributeName
	ModelSlice []*attribute_name.AttributeName
}
type AttributeNameResource struct {
	ID          uint64         `json:"id"`
	State       uint8          `json:"state"`
	IsPublic    uint8          `json:"public"`
	Title       string         `json:"title"`
	Abbr        string         `json:"abbr"`
	Genre       uint8          `json:"genre"`
	Sort        uint64         `json:"sort"`
	Search      string         `json:"search"`
	Description string         `json:"desc"`
	Categories  []*CategoryHas `json:"categories"`
	CreatedAt   string         `json:"created_at"`
	UpdatedAt   string         `json:"updated_at"`
}

func (p *AttributeName) ShowResource() (show *AttributeNameResource) {
	optimus := optimusPkg.NewOptimus()
	show.ID = optimus.Encode(p.Model.ID)
	show.State = p.Model.State
	show.IsPublic = p.Model.IsPublic
	show.Title = p.Model.Title
	show.Description = p.Model.Description
	show.Sort = p.Model.Sort
	show.Abbr = p.Model.Abbr
	show.Search = p.Model.Search
	show.Genre = p.Model.Genre
	show.Categories = CategoryHasResource(p.Model.Category)
	show.CreatedAt = helpers.TimeFormat(p.Model.CreatedAt, "second")
	show.UpdatedAt = helpers.TimeFormat(p.Model.UpdatedAt, "second")
	return
}
func (p *AttributeName) IndexResource() (index []*AttributeNameResource) {
	optimus := optimusPkg.NewOptimus()
	for _, model := range p.ModelSlice {
		index = append(index, &AttributeNameResource{
			ID:          optimus.Encode(model.ID),
			State:       model.State,
			IsPublic:    model.IsPublic,
			Title:       model.Title,
			Description: model.Description,
			Categories:  CategoryHasResource(model.Category),
			Sort:        model.Sort,
			Abbr:        model.Abbr,
			Search:      model.Search,
			Genre:       model.Genre,
			CreatedAt:   helpers.TimeFormat(model.CreatedAt, "second"),
			UpdatedAt:   helpers.TimeFormat(model.UpdatedAt, "second"),
		})
	}
	return
}

type AttributeNameSelect struct {
	ID    uint64 `json:"value"`
	Title string `json:"title"`
}

func (p *AttributeName) InitialResource() (sel []*AttributeNameSelect) {
	optimus := optimusPkg.NewOptimus()

	for _, model := range p.ModelSlice {
		sel = append(sel, &AttributeNameSelect{
			ID:    optimus.Encode(model.ID),
			Title: model.Title,
		})
	}
	return
}
