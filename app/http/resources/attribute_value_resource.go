package resources

import (
	"item-server/app/models/attribute_value"
	"item-server/pkg/helpers"
	optimusPkg "item-server/pkg/optimus"
)

type AttributeValue struct {
	Model      *attribute_value.AttributeValue
	ModelSlice []*attribute_value.AttributeValue
}

func (p *AttributeValue) ShowResource() map[string]any {
	optimus := optimusPkg.NewOptimus()
	return map[string]any{
		"id":             optimus.Encode(p.Model.ID),
		"state":          p.Model.State,
		"title":          p.Model.Title,
		"description":    p.Model.Description,
		"attribute_name": optimus.Encode(p.Model.AttributeNameId),
		"sort":           p.Model.Sort,
		"abbr":           p.Model.Abbr,
		"search":         p.Model.Search,
		"created_at":     helpers.TimeFormat(p.Model.CreatedAt, "second"),
		"updated_at":     helpers.TimeFormat(p.Model.UpdatedAt, "second"),
	}
}
func (p *AttributeValue) IndexResource() []any {
	optimus := optimusPkg.NewOptimus()
	s := make([]any, 0)
	for _, model := range p.ModelSlice {
		s = append(s, map[string]any{
			"id":             optimus.Encode(model.ID),
			"state":          model.State,
			"title":          model.Title,
			"description":    model.Description,
			"attribute_name": optimus.Encode(model.AttributeNameId),
			"sort":           model.Sort,
			"abbr":           model.Abbr,
			"search":         model.Search,
			"attribute_abbr": model.AttributeName.Title,
			"created_at":     helpers.TimeFormat(model.CreatedAt, "second"),
			"updated_at":     helpers.TimeFormat(model.UpdatedAt, "second"),
		})
	}
	return s
}
