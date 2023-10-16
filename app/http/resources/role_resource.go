package resources

import (
	"item-server/app/models/permission"
	"item-server/app/models/role"
	"item-server/pkg/helpers"
	optimusPkg "item-server/pkg/optimus"
)

type Role struct {
	Model      *role.Role
	ModelSlice []role.Role
	PerSlice   []permission.Permission
}

func (r *Role) ShowResource() map[string]any {
	optimus := optimusPkg.NewOptimus()
	return map[string]any{
		"id":          optimus.Encode(r.Model.ID),
		"state":       r.Model.State,
		"name":        r.Model.Name,
		"title":       r.Model.Title,
		"guard":       r.Model.GuardName,
		"description": r.Model.Description,
		"permissions": permissionCollection(r.Model.Permissions),
		"permission":  permissionSlice(r.Model.Permissions),
		"created_at":  helpers.TimeFormat(r.Model.CreatedAt, "second"),
		"updated_at":  helpers.TimeFormat(r.Model.UpdatedAt, "second"),
	}
}
func (r *Role) IndexResource() []any {
	optimus := optimusPkg.NewOptimus()
	s := make([]any, 0)
	for _, model := range r.ModelSlice {
		s = append(s, map[string]any{
			"id":          optimus.Encode(model.ID),
			"state":       model.State,
			"name":        model.Name,
			"title":       model.Title,
			"guard":       model.GuardName,
			"description": model.Description,
			"permissions": permissionCollection(model.Permissions),
			"permission":  permissionSlice(model.Permissions),
			"created_at":  helpers.TimeFormat(model.CreatedAt, "second"),
			"updated_at":  helpers.TimeFormat(model.UpdatedAt, "second"),
		})
	}
	return s
}
func (r *Role) InitialPerSliceResource() []any {
	optimus := optimusPkg.NewOptimus()
	s := make([]any, 0)
	for _, model := range r.PerSlice {
		s = append(s, map[string]any{
			"value": optimus.Encode(model.ID),
			"label": model.Title,
		})
	}
	return s
}

func permissionCollection(r []permission.Permission) []any {
	optimus := optimusPkg.NewOptimus()
	s := make([]interface{}, 0)
	for _, v := range r {
		m := map[string]interface{}{
			"id":    optimus.Encode(v.ID),
			"name":  v.Name,
			"title": v.Title,
			"guard": v.GuardName,
		}
		s = append(s, m)
	}
	return s
}

func permissionSlice(p []permission.Permission) []uint64 {
	optimus := optimusPkg.NewOptimus()
	i := make([]uint64, 0)
	for _, v := range p {
		i = append(i, optimus.Encode(v.ID))
	}
	return i
}
