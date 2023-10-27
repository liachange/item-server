package resources

import (
	"item-server/app/models/category"
	"item-server/app/models/unit"
	"item-server/pkg/helpers"
	optimusPkg "item-server/pkg/optimus"
)

type Unit struct {
	Model      *unit.Unit
	ModelSlice []unit.Unit
}

func (p *Unit) ShowResource() map[string]any {
	optimus := optimusPkg.NewOptimus()
	return map[string]any{
		"id":          optimus.Encode(p.Model.ID),
		"state":       p.Model.State,
		"public":      p.Model.IsPublic,
		"title":       p.Model.Title,
		"description": p.Model.Description,
		"sort":        p.Model.Sort,
		"categories":  category.OptId(p.Model.Category),
		"created_at":  helpers.TimeFormat(p.Model.CreatedAt, "second"),
		"updated_at":  helpers.TimeFormat(p.Model.UpdatedAt, "second"),
	}
}
func (p *Unit) IndexResource() []any {
	optimus := optimusPkg.NewOptimus()
	s := make([]any, 0)
	for _, model := range p.ModelSlice {
		s = append(s, map[string]any{
			"id":          optimus.Encode(model.ID),
			"state":       model.State,
			"public":      model.IsPublic,
			"title":       model.Title,
			"description": model.Description,
			"categories":  category.OptId(model.Category),
			"sort":        model.Sort,
			"created_at":  helpers.TimeFormat(model.CreatedAt, "second"),
			"updated_at":  helpers.TimeFormat(model.UpdatedAt, "second"),
		})
	}
	return s
}
func (p *Unit) InitialResource() []any {
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
