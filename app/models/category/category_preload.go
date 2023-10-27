package category

import (
	optimusPkg "item-server/pkg/optimus"
)

type Many struct {
	ID    uint64 `gorm:"column:id;primaryKey;autoIncrement;" json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Abbr  string `json:"abbr,omitempty"`
}

// TableName 会将 AttributeNameMany 的表名重写为 `attribute_names`
func (*Many) TableName() string {
	return "categories"
}

func OptId(r []*Many) []any {
	optimus := optimusPkg.NewOptimus()
	s := make([]interface{}, 0)
	for _, v := range r {
		m := map[string]interface{}{
			"id":    optimus.Encode(v.ID),
			"title": v.Title,
		}
		s = append(s, m)
	}
	return s
}
