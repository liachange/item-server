package brand

import (
	"errors"
	"gorm.io/gorm"
	"item-server/app/models"
	"item-server/app/models/category"
	"item-server/pkg/database"
	"item-server/pkg/helpers"
)

type Brand struct {
	models.BaseModel

	State       uint8            `json:"state,omitempty"`
	Title       string           `json:"title,omitempty"`
	Description string           `json:"description,omitempty"`
	IconUrl     string           `json:"icon_url,omitempty"`
	Sort        uint64           `json:"sort,omitempty"`
	IsPublic    uint8            `json:"is_public,omitempty"`
	Category    []*category.Many `gorm:"many2many:category_brands;foreignKey:ID;joinForeignKey:brand_id;references:ID;joinReferences:category_id;" json:"categories"`
	models.DeletedAt
	models.CommonTimestampsField
}

type CategoryBrand struct {
	CategoryId uint64 `json:"category_id,omitempty"`
	BrandId    uint64 `json:"brand_id,omitempty"`
}

func (brand *Brand) CreateMany(permKey []uint64) (id uint64) {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := (tx.Create(&brand)).Error; err != nil {
			return err
		}
		var categoryBrands []*CategoryBrand
		if brand.ID > 0 {
			if len(permKey) != 0 {
				for _, v := range permKey {
					row := &CategoryBrand{
						CategoryId: v,
						BrandId:    brand.ID,
					}
					categoryBrands = append(categoryBrands, row)
				}
				if err := (tx.Create(&categoryBrands)).Error; err != nil {
					return err
				}
			}
		} else {
			return errors.New("参数异常...")
		}
		return nil
	})
	if err == nil {
		id = brand.ID
	}
	return id
}

func (brand *Brand) SaveMany(sel []string, permKey []uint64) (id uint64) {

	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := (tx.Select(sel).Save(&brand)).Error; err != nil {
			return err
		}
		if err := (tx.Where("brand_id=?", brand.ID).Delete(&CategoryBrand{})).Error; err != nil {
			return err
		}
		var categoryBrands []*CategoryBrand
		if brand.ID > 0 {
			if !helpers.Empty(permKey) {
				for _, v := range permKey {
					row := &CategoryBrand{
						CategoryId: v,
						BrandId:    brand.ID,
					}
					categoryBrands = append(categoryBrands, row)
				}
				if err := (tx.Create(&categoryBrands)).Error; err != nil {
					return err
				}
			}

		} else {
			return errors.New("参数异常...")
		}
		return nil
	})
	if err == nil {
		id = brand.ID
	}
	return id
}

func (brand *Brand) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&brand)
	return result.RowsAffected
}
