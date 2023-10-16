package resources

import (
	"item-server/app/models/permission"
	"item-server/pkg/helpers"
	optimusPkg "item-server/pkg/optimus"
)

type Permission struct {
	Model      *permission.Permission
	ModelSlice []permission.Permission
}

func (p *Permission) ShowResource() map[string]any {
	optimus := optimusPkg.NewOptimus()
	parent := p.Model.ParentID
	if parent > 0 {
		parent = optimus.Encode(p.Model.ParentID)
	}
	return map[string]any{
		"id":          optimus.Encode(p.Model.ID),
		"state":       p.Model.State,
		"name":        p.Model.Name,
		"title":       p.Model.Title,
		"guard":       p.Model.GuardName,
		"type":        p.Model.Type,
		"parent":      parent,
		"sort":        p.Model.Sort,
		"icon":        p.Model.Icon,
		"description": p.Model.Description,
		"created_at":  helpers.TimeFormat(p.Model.CreatedAt, "second"),
		"updated_at":  helpers.TimeFormat(p.Model.UpdatedAt, "second"),
	}
}
func (p *Permission) IndexResource() []any {
	optimus := optimusPkg.NewOptimus()
	s := make([]any, 0)
	for _, model := range p.ModelSlice {
		parent := model.ParentID
		if parent > 0 {
			parent = optimus.Encode(model.ParentID)
		}
		s = append(s, map[string]any{
			"id":          optimus.Encode(model.ID),
			"state":       model.State,
			"name":        model.Name,
			"title":       model.Title,
			"guard":       model.GuardName,
			"type":        model.Type,
			"parent":      parent,
			"sort":        model.Sort,
			"icon":        model.Icon,
			"description": model.Description,
			"created_at":  helpers.TimeFormat(model.CreatedAt, "second"),
			"updated_at":  helpers.TimeFormat(model.UpdatedAt, "second"),
		})
	}
	return s
}
func (p *Permission) InitialResource() []any {
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
