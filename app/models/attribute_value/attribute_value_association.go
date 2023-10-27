package attribute_value

type AttributeNameOne struct {
	ID    uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Abbr  string `json:"abbr,omitempty"`
}

// TableName 会将 RoleMany 的表名重写为 `attribute_names`
func (*AttributeNameOne) TableName() string {
	return "attribute_names"
}
