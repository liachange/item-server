package unit

import (
	"errors"
	"gorm.io/gorm"
	"item-server/app/models"
	"item-server/app/models/category"
	"item-server/pkg/database"
	"item-server/pkg/helpers"
)

type Unit struct {
	models.BaseModel

	State       uint8            `json:"state,omitempty"`
	Title       string           `json:"title,omitempty"`
	Description string           `json:"description,omitempty"`
	Sort        uint64           `json:"sort,omitempty"`
	Abbr        string           `json:"abbr,omitempty"`
	IsPublic    uint8            `json:"is_public,omitempty"`
	Category    []*category.Many `gorm:"many2many:category_units;foreignKey:ID;joinForeignKey:unit_id;references:ID;joinReferences:category_id;" json:"categories"`

	models.DeletedAt
	models.CommonTimestampsField
}
type CategoryUnit struct {
	CategoryId uint64 `json:"category_id,omitempty"`
	UnitId     uint64 `json:"unit_id,omitempty"`
}

func (unit *Unit) Create() {
	database.DB.Create(&unit)
}

func (unit *Unit) CreateMany(permKey []uint64) (id uint64) {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := (tx.Create(&unit)).Error; err != nil {
			return err
		}
		var categoryUnits []*CategoryUnit
		if unit.ID > 0 {
			if !helpers.Empty(permKey) {
				for _, v := range permKey {
					row := &CategoryUnit{
						CategoryId: v,
						UnitId:     unit.ID,
					}
					categoryUnits = append(categoryUnits, row)
				}
				if err := (tx.Create(&categoryUnits)).Error; err != nil {
					return err
				}
			}
		} else {
			return errors.New("参数异常...")
		}
		return nil
	})
	if err == nil {
		id = unit.ID
	}
	return id
}
func (unit *Unit) SaveMany(fieldSelect []string, permKey []uint64) (id uint64) {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := (tx.Select(fieldSelect).Save(&unit)).Error; err != nil {
			return err
		}
		if err := (tx.Where("unit_id=?", unit.ID).Delete(&CategoryUnit{})).Error; err != nil {
			return err
		}
		var categoryUnits []*CategoryUnit
		if unit.ID > 0 {
			if !helpers.Empty(permKey) {
				for _, v := range permKey {
					row := &CategoryUnit{
						CategoryId: v,
						UnitId:     unit.ID,
					}
					categoryUnits = append(categoryUnits, row)
				}
				if err := (tx.Create(&categoryUnits)).Error; err != nil {
					return err
				}
			}

		} else {
			return errors.New("参数异常...")
		}
		return nil
	})
	if err == nil {
		id = unit.ID
	}
	return id
}

func (unit *Unit) Save(fieldSelect any) (rowsAffected int64) {
	result := database.DB.Select(helpers.ReqSelect(fieldSelect)).Save(&unit)
	return result.RowsAffected
}

func (unit *Unit) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&unit)
	return result.RowsAffected
}
