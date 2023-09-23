// Package role_has_permission 模型
package role_has_permission

type RoleHasPermission struct {
	RoleID       uint64 `json:"role_id,omitempty"`
	PermissionID uint64 `json:"permission_id,omitempty"`
}
