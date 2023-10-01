package v1

import (
	"item-server/app/models/role"
	"item-server/app/models/user"
	"item-server/pkg/helpers"
)

func IndexResource(users []user.User) (s []interface{}) {
	for _, v := range users {
		m := map[string]interface{}{
			"id":          v.ID,
			"state":       v.State,
			"name":        v.Name,
			"phone":       v.Phone,
			"email":       v.Email,
			"avatar_icon": v.Avatar,
			"nickname":    v.Nickname,
			"roles":       roleCollection(v.Role),
			"created_at":  helpers.TimeFormat(v.CreatedAt, "second"),
			"updated_at":  helpers.TimeFormat(v.UpdatedAt, "second"),
		}
		s = append(s, m)
	}
	return
}

func ShowResource(user user.User) map[string]interface{} {
	return map[string]interface{}{
		"id":         user.ID,
		"name":       user.Name,
		"roles":      roleCollection(user.Role),
		"created_at": helpers.TimeFormat(user.CreatedAt, "second"),
		"updated_at": helpers.TimeFormat(user.UpdatedAt, "second"),
	}
}

func roleCollection(r []role.Role) (s []interface{}) {
	for _, v := range r {
		m := map[string]interface{}{
			"id":    v.ID,
			"name":  v.Name,
			"guard": v.GuardName,
		}
		s = append(s, m)
	}
	return
}
