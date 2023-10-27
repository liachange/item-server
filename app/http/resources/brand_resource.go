package resources

import (
	"item-server/app/models/brand"
	"item-server/app/models/category"
	"item-server/pkg/helpers"
	optimusPkg "item-server/pkg/optimus"
)

type Brand struct {
	Model      *brand.Brand
	ModelSlice []brand.Brand
}

func (p *Brand) ShowResource() map[string]any {
	optimus := optimusPkg.NewOptimus()

	return map[string]any{
		"id":          optimus.Encode(p.Model.ID),
		"state":       p.Model.State,
		"title":       p.Model.Title,
		"description": p.Model.Description,
		"icon":        p.Model.IconUrl,
		"sort":        p.Model.Sort,
		"categories":  category.OptId(p.Model.Category),
		"created_at":  helpers.TimeFormat(p.Model.CreatedAt, "second"),
		"updated_at":  helpers.TimeFormat(p.Model.UpdatedAt, "second"),
	}
}
func (p *Brand) IndexResource() []any {
	optimus := optimusPkg.NewOptimus()
	s := make([]any, 0)
	for _, model := range p.ModelSlice {

		s = append(s, map[string]any{
			"id":          optimus.Encode(model.ID),
			"state":       model.State,
			"title":       model.Title,
			"description": model.Description,
			"icon":        model.IconUrl,
			"sort":        model.Sort,
			"categories":  category.OptId(model.Category),
			"created_at":  helpers.TimeFormat(model.CreatedAt, "second"),
			"updated_at":  helpers.TimeFormat(model.UpdatedAt, "second"),
		})
	}
	return s
}
func (p *Brand) InitialResource() []any {
	optimus := optimusPkg.NewOptimus()
	s := make([]any, 0)
	for _, model := range p.ModelSlice {
		s = append(s, map[string]any{
			"value": optimus.Encode(model.ID),
			"label": model.Title,
		})
	}
	return s
}
