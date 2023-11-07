package attribute_name

import (
	"errors"
	"gorm.io/gorm"
	"item-server/app/models"
	"item-server/app/models/category"
	"item-server/pkg/database"
	"item-server/pkg/helpers"
)

const (
	sku = iota + 1
	item
	other
)

type AttributeName struct {
	models.BaseModel

	State       uint8            `json:"state,omitempty"`
	IsPublic    uint8            `json:"is_public,omitempty"`
	Genre       uint8            `json:"genre,omitempty"`
	Title       string           `json:"title,omitempty"`
	Description string           `json:"description,omitempty"`
	Sort        uint64           `json:"sort,omitempty"`
	Abbr        string           `json:"abbr,omitempty"`
	Search      string           `json:"search,omitempty"`
	Category    []*category.Many `gorm:"many2many:category_attribute_names;foreignKey:ID;joinForeignKey:attribute_name_id;references:ID;joinReferences:category_id;" json:"categories"`
	models.DeletedAt
	models.CommonTimestampsField
}
type CategoryAttributeName struct {
	CategoryId      uint64 `json:"category_id,omitempty"`
	AttributeNameId uint64 `json:"attribute_name_id,omitempty"`
}

func (attributeName *AttributeName) CreateMany(permKey []uint64) (id uint64) {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := (tx.Create(&attributeName)).Error; err != nil {
			return err
		}
		var categoryAttributeNames []*CategoryAttributeName
		if attributeName.ID > 0 {
			if !helpers.Empty(permKey) {
				for _, v := range permKey {
					row := &CategoryAttributeName{
						CategoryId:      v,
						AttributeNameId: attributeName.ID,
					}
					categoryAttributeNames = append(categoryAttributeNames, row)
				}
				if err := (tx.Create(&categoryAttributeNames)).Error; err != nil {
					return err
				}
			}
		} else {
			return errors.New("参数异常...")
		}
		return nil
	})
	if err == nil {
		id = attributeName.ID
	}
	return id
}

func AttributeNameAll() (attributeName []*AttributeName) {
	database.DB.Select("id", "title", "abbr").Order("sort asc").Find(&attributeName)
	return
}

func (attributeName *AttributeName) SaveMany(sel []string, permKey []uint64) (id uint64) {

	err := database.DB.Transaction(func(tx *gorm.DB) error {

		if result := tx.Select(sel).Save(&attributeName); result.Error != nil {
			return result.Error
		}

		if err := (tx.Where("attribute_name_id=?", attributeName.ID).Delete(&CategoryAttributeName{})).Error; err != nil {
			return err
		}
		var categoryAttributeNames []*CategoryAttributeName
		if attributeName.ID > 0 {
			if len(permKey) != 0 {
				for _, v := range permKey {
					row := &CategoryAttributeName{
						CategoryId:      v,
						AttributeNameId: attributeName.ID,
					}
					categoryAttributeNames = append(categoryAttributeNames, row)
				}
				if err := (tx.Create(&categoryAttributeNames)).Error; err != nil {
					return err
				}
			}

		} else {
			return errors.New("参数异常...")
		}
		return nil
	})
	if err == nil {
		id = attributeName.ID
	}
	return id
}

func (attributeName *AttributeName) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&attributeName)
	return result.RowsAffected
}

func InitGenre() []map[string]any {
	return []map[string]any{
		{
			"value": sku,
			"label": "规格",
		},
		{
			"value": item,
			"label": "属性",
		},
		{
			"value": other,
			"label": "其他",
		},
	}
}
