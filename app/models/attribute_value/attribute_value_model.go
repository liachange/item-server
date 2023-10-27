package attribute_value

import (
	"item-server/app/models"
	"item-server/pkg/database"
	"item-server/pkg/helpers"
)

type AttributeValue struct {
	models.BaseModel

	State           uint8            `json:"state,omitempty"`
	AttributeNameId uint64           `json:"attribute_name,omitempty"`
	Title           string           `json:"title,omitempty"`
	Description     string           `json:"description,omitempty"`
	Sort            uint64           `json:"sort,omitempty"`
	Abbr            string           `json:"abbr,omitempty"`
	Search          string           `json:"search,omitempty"`
	AttributeName   AttributeNameOne `gorm:"foreignKey:attribute_name_id" json:"att_name,omitempty"`
	models.DeletedAt

	models.CommonTimestampsField
}

func (attributeValue *AttributeValue) Create() {
	database.DB.Create(&attributeValue)
}

func (attributeValue *AttributeValue) Save(fieldSelect any) (rowsAffected int64) {
	result := database.DB.Select(helpers.ReqSelect(fieldSelect)).Save(&attributeValue)
	return result.RowsAffected
}

func (attributeValue *AttributeValue) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&attributeValue)
	return result.RowsAffected
}
