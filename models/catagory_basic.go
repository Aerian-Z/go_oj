package models

import "gorm.io/gorm"

type CategoryBasic struct {
	gorm.Model
	Identity string `gorm:"column:identity; type:varchar(36); json:"identity"` // unique identification of the category table
	Name     string `gorm:"column:name; type:varchar(255); json:"name"`
	ParentId int    `gorm:"column:parent_id; type:int; json:"parent_id"`
}

func (table *CategoryBasic) TableName() string {
	return "category_basic"
}

func GetCategoryList(keyword string) *gorm.DB {
	return DB.Model(new(CategoryBasic)).Where("name LIKE ? ", "%"+keyword+"%")
}
