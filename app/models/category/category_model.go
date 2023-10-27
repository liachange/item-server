package category

import (
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"item-server/app/models"
	"item-server/pkg/database"
	"item-server/pkg/helpers"
)

type Category struct {
	models.BaseModel

	State       uint8  `json:"state,omitempty"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	IconUrl     string `json:"icon_url,omitempty"`
	Sort        uint64 `json:"sort,omitempty"`
	ParentId    uint64 `json:"parent,omitempty"`
	Level       uint8  `json:"level,omitempty"`
	LevelTree   string `json:"level_tree,omitempty"`

	models.DeletedAt
	models.CommonTimestampsField
}

func (category *Category) Create() (id uint64) {
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := (tx.Create(&category)).Error; err != nil {
			return err
		}
		var level uint8 = 1
		tree := "." + cast.ToString(category.ID) + "."
		if category.ParentId > 0 {
			var parentCategory Category
			tx.Select("id", "level", "level_tree").First(&parentCategory, category.ParentId)
			level = parentCategory.Level + 1
			tree = parentCategory.LevelTree + cast.ToString(category.ID) + "."
		}
		tx.Where("id=?", category.ID).Updates(Category{Level: level, LevelTree: tree})
		return nil
	})
	if err == nil {
		id = category.ID
	}
	return id
}

func (category *Category) Save(fieldSelect any) (rowsAffected int64) {
	var level uint8 = 1
	tree := "." + cast.ToString(category.ID) + "."
	if category.ParentId > 0 {
		var parentCategory Category
		database.DB.Select("id", "level", "level_tree").First(&parentCategory, category.ParentId)
		level = parentCategory.Level + 1
		tree = parentCategory.LevelTree + cast.ToString(category.ID) + "."
	}
	category.Level = level
	category.LevelTree = tree
	//追加更新字段
	selectStr := helpers.ReqSelect(fieldSelect)
	selectStr = append(selectStr, "level", "level_tree")
	result := database.DB.Select(selectStr).Save(&category)
	return result.RowsAffected
}

func (category *Category) Delete() (rowsAffected int64) {
	result := database.DB.Model(&category).Updates(map[string]interface{}{
		"deleted_at": helpers.TimeNow(),
		"title":      category.Title + "_remove",
	})
	return result.RowsAffected
}
