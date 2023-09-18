package user

import "item-server/app/models"

// User 用户模型
type User struct {
	models.BaseModel

	Name     string `json:"name,omitempty"`
	Email    string `json:"-"`
	Phone    string `json:"-"`
	Password string `json:"-"`

	Avatar   string `json:"avatar_icon,omitempty"`
	Nickname string `json:"nickname,omitempty"`

	models.CommonTimestampsField
}
