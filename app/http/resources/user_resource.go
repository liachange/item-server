package resources

import (
	"item-server/app/models/role"
	"item-server/app/models/user"
	"item-server/pkg/helpers"
	optimusPkg "item-server/pkg/optimus"
)

type User struct {
	Model      *user.User
	ModelSlice []user.User
	RoleSlice  []role.Role
	Menu       *user.UserMany
}
type UserMenu struct {
	ID       uint64     `json:"id,omitempty"`
	Name     string     `json:"name,omitempty"`
	Title    string     `json:"title,omitempty"`
	ParentID uint64     `json:"parent_id"`
	Icon     string     `json:"icon"`
	Sort     uint64     `json:"sort,omitempty"`
	Children []UserMenu `json:"children"`
}

func (u *User) ShowResource() map[string]any {
	optimus := optimusPkg.NewOptimus()
	return map[string]any{
		"id":          optimus.Encode(u.Model.ID),
		"state":       u.Model.State,
		"name":        u.Model.Name,
		"phone":       u.Model.Phone,
		"email":       u.Model.Email,
		"avatar_icon": u.Model.Avatar,
		"nickname":    u.Model.Nickname,
		"roles":       roleCollection(u.Model.Role),
		"role":        roleSlice(u.Model.Role),
		"created_at":  helpers.TimeFormat(u.Model.CreatedAt, "second"),
		"updated_at":  helpers.TimeFormat(u.Model.UpdatedAt, "second"),
	}
}
func (u *User) IndexResource() []any {
	optimus := optimusPkg.NewOptimus()
	s := make([]any, 0)
	for _, model := range u.ModelSlice {
		s = append(s, map[string]any{
			"id":          optimus.Encode(model.ID),
			"state":       model.State,
			"name":        model.Name,
			"phone":       model.Phone,
			"email":       model.Email,
			"avatar_icon": model.Avatar,
			"nickname":    model.Nickname,
			"roles":       roleCollection(model.Role),
			"role":        roleSlice(model.Role),
			"created_at":  helpers.TimeFormat(model.CreatedAt, "second"),
			"updated_at":  helpers.TimeFormat(model.UpdatedAt, "second"),
		})
	}
	return s
}
func (u *User) InitialRoleResource() []any {
	optimus := optimusPkg.NewOptimus()
	s := make([]any, 0)
	for _, model := range u.RoleSlice {
		s = append(s, map[string]any{
			"value": optimus.Encode(model.ID),
			"label": model.Title,
		})
	}
	return s
}

func (u *User) MenuResource() []UserMenu {
	var list []UserMenu
	for _, r := range u.Menu.Roles {
		for _, p := range r.Permissions {
			list = append(list, UserMenu{
				ID:       p.ID,
				Name:     p.Name,
				Title:    p.Title,
				ParentID: p.ParentID,
				Icon:     p.Icon,
				Sort:     p.Sort,
			})

		}
	}
	return menu(list, 0)
}

func roleCollection(r []role.Role) []any {
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

func roleSlice(p []role.Role) []uint64 {
	optimus := optimusPkg.NewOptimus()
	i := make([]uint64, 0)
	for _, v := range p {
		i = append(i, optimus.Encode(v.ID))
	}
	return i
}

func menu(list []UserMenu, parent uint64) []UserMenu {
	var menuList []UserMenu
	for _, r := range list {
		if r.ParentID == parent {
			r.Children = menu(list, r.ID)
			menuList = append(menuList, r)
		}
	}
	return menuList
}
