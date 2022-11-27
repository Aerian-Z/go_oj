package models

import "gorm.io/gorm"

type ProblemCategory struct {
	gorm.Model
	ProblemId     string         `gorm:"column:problem_id; type:varchar(36); json:"problem_id"`   // ID of the problem table
	CategoryId    string         `gorm:"column:category_id; type:varchar(36); json:"category_id"` // ID of the category table
	CategoryBasic *CategoryBasic `gorm:"foreignKey:id; references:category_id; json:"category_basic"`
}

func (table *ProblemCategory) TableName() string {
	return "problem_category"
}
