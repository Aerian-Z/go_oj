package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Identity string `gorm:"column:identity; type:varchar(36); json:"identity"` // unique identification of the user table
	Username string `gorm:"column:username; type:varchar(255); json:"username"`
	Password string `gorm:"column:password; type:varchar(255); json:"password"`
	Phone    string `gorm:"column:phone; type:varchar(255); json:"phone"`
	Mail     string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserDetail(userIdentity string) *gorm.DB {
	return DB.Model(new(UserBasic)).Omit("password").Where("identity = ?", userIdentity)
}
