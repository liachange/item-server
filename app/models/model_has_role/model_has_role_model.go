// Package model_has_role 模型
package model_has_role

type ModelHasRole struct {
	RoleID    uint64 `json:"role_id,omitempty"`
	ModelType string `json:"model_type,omitempty"`
	ModelID   uint64 `json:"model_id,omitempty"`
}
