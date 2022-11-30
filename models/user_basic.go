package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Identity  string `gorm:"column:identity; type:varchar(36); json:"identity"` // unique identification of the user table
	Username  string `gorm:"column:username; type:varchar(255); json:"username"`
	Password  string `gorm:"column:password; type:varchar(255); json:"password"`
	Phone     string `gorm:"column:phone; type:varchar(255); json:"phone"`
	Email     string `gorm:"column:email; type:varchar(255); json:"email"`
	PassNum   int    `gorm:"column:pass_num; type:int; json:"pass_num"`
	SubmitNum int64  `gorm:"column:submit_num; type:bigint; json:"submit_num"`
	IsAdmin   int    `gorm:"column:is_admin; type:int; json:"is_admin"` // 0-not admin, 1-admin
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserDetail(userIdentity string) *gorm.DB {
	return DB.Model(new(UserBasic)).Omit("password").Where("identity = ?", userIdentity)
}
