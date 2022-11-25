package models

import "gorm.io/gorm"

type Problem struct {
	gorm.Model
	Identidy   string `gorm:"column:identify; type:varchar(36); json:"identity"`        // unique identification of the problem table
	CatagoryID string `gorm:"column:catagory_id; type:varchar(255); json:"catagory_id"` // catagory id, separated by commas
	Title      string `gorm:"column:title; type:varchar(255); json:"title"`             // title of the article
	Content    string `gorm:"column: content; type:longtext; json:"content"`            // content of the article
	MaxRuntime int    `gorm:"column:max_runtime; type:int; json:"max_runtime"`          // max runtime of the problem
	MaxMenory  int    `gorm:"column:max_menory; type:int; json:"max_menory"`            // max menory of the problem
}

func (table *Problem) TableName() string {
	return "problem"
}
