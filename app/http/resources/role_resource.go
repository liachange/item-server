package resources

import (
	"item-server/app/models/permission"
	"item-server/app/models/role"
	"item-server/pkg/helpers"
)

func RoleIndexResource(models []role.Role) []interface{} {
	s := make([]interface{}, 0)
	for _, model := range models {
		s = append(s, map[string]interface{}{
			"id":          model.ID,
			"state":       model.State,
			"name":        model.Name,
			"title":       model.Title,
			"guard":       model.GuardName,
			"description": model.Description,
			"permissions": permissionCollection(model.Permissions),
			"created_at":  helpers.TimeFormat(model.CreatedAt, "second"),
			"updated_at":  helpers.TimeFormat(model.UpdatedAt, "second"),
		})
	}
	return s
}

func RoleShowResource(model role.Role) map[string]interface{} {
	return map[string]interface{}{
		"id":          model.ID,
		"state":       model.State,
		"name":        model.Name,
		"title":       model.Title,
		"guard":       model.GuardName,
		"description": model.Description,
		"permissions": permissionCollection(model.Permissions),
		"created_at":  helpers.TimeFormat(model.CreatedAt, "second"),
		"updated_at":  helpers.TimeFormat(model.UpdatedAt, "second"),
	}
}
func permissionCollection(r []permission.Permission) []any {
	s := make([]interface{}, 0)
	for _, v := range r {
		m := map[string]interface{}{
			"id":    v.ID,
			"name":  v.Name,
			"guard": v.GuardName,
		}
		s = append(s, m)
	}
	return s
}
