package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Identidy string `gorm:"column:identify; type:varchar(36); json:"identity"` // unique identification of the user table
	Name     string `gorm:"column:name; type:varchar(255); json:"name"`
	Password string `gorm:"column:password; type:varchar(255); json:"password"`
	Phone    string `gorm:"column:phone; type:varchar(255); json:"phone"`
	Mail     string
}

func (table *User) TableName() string {
	return "user"
}
