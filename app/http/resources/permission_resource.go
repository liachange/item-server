package resources

import (
	"item-server/app/models/permission"
	"item-server/pkg/helpers"
)

func PermissionIndexResource(models []permission.Permission) []interface{} {
	s := make([]interface{}, 0)
	for _, model := range models {
		s = append(s, map[string]interface{}{
			"id":          model.ID,
			"state":       model.State,
			"name":        model.Name,
			"title":       model.Title,
			"guard":       model.GuardName,
			"type":        model.Type,
			"parent":      model.ParentID,
			"sort":        model.Sort,
			"icon":        model.Icon,
			"description": model.Description,
			"created_at":  helpers.TimeFormat(model.CreatedAt, "second"),
			"updated_at":  helpers.TimeFormat(model.UpdatedAt, "second"),
		})
	}
	return s
}

func PermissionShowResource(model permission.Permission) map[string]interface{} {
	return map[string]interface{}{
		"id":          model.ID,
		"state":       model.State,
		"name":        model.Name,
		"title":       model.Title,
		"guard":       model.GuardName,
		"type":        model.Type,
		"parent":      model.ParentID,
		"sort":        model.Sort,
		"icon":        model.Icon,
		"description": model.Description,
		"created_at":  helpers.TimeFormat(model.CreatedAt, "second"),
		"updated_at":  helpers.TimeFormat(model.UpdatedAt, "second"),
	}
}
