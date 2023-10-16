package user

// UserMany  many2many
// foreignKey 主表标识 对应 joinForeignKey 中间表 标识
// references 引用表标识 对应joinReferences 中间表 标识
type UserMany struct {
	ID    uint64      `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
	State uint8       `json:"state,omitempty"`
	Name  string      `json:"name,omitempty"`
	Roles []*RoleMany `gorm:"many2many:model_has_roles;foreignKey:ID;joinForeignKey:model_id;references:ID;joinReferences:role_id;" json:"roles"`
}

// TableName 会将 UserMany 的表名重写为 `users`
func (*UserMany) TableName() string {
	return "users"
}

type RoleMany struct {
	ID          uint64           `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
	State       uint8            `json:"state,omitempty"`
	Permissions []PermissionMany `gorm:"many2many:role_has_permissions;foreignKey:ID;joinForeignKey:role_id;references:ID;joinReferences:permission_id;" json:"permissions"`
}

// TableName 会将 RoleMany 的表名重写为 `roles`
func (*RoleMany) TableName() string {
	return "roles"
}

type PermissionMany struct {
	ID       uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
	State    uint8  `json:"state,omitempty"`
	Name     string `json:"name,omitempty"`
	Title    string `json:"title,omitempty"`
	ParentID uint64 `json:"parent"`
	Sort     uint64 `json:"sort,omitempty"`
	Icon     string `json:"icon"`
}

// TableName 会将 PermissionMany 的表名重写为 `permissions`
func (*PermissionMany) TableName() string {
	return "permissions"
}
