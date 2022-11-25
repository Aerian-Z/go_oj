package models

import "gorm.io/gorm"

type Catagory struct {
	gorm.Model
	Identidy string `gorm:"column:identify; type:varchar(36); json:"identity"` // unique identification of the catagory table
	Name     string `gorm:"column:name; type:varchar(255); json:"name"`
	ParentId int    `gorm:"column:parent_id; type:int; json:"parent_id"`
}

func (table *Catagory) TableName() string {
	return "catagory"
}
