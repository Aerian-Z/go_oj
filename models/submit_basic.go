package models

import "gorm.io/gorm"

type SubmitBasic struct {
	gorm.Model
	Identity        string        `gorm:"column:identity; type:varchar(36); json:"identity"`                 // unique identification of the submit table
	ProblemIdentity string        `gorm:"column:problem_identity; type:varchar(36); json:"problem_identity"` // unique identification of the problem table
	ProblemBasic    *ProblemBasic `gorm:"foreignKey:identity; references:problem_identity; json:"problem_basic"`
	UserIdentity    string        `gorm:"column:user_identity; type:varchar(36); json:"user_identity"` // unique identification of the user table
	UserBasic       *UserBasic    `gorm:"foreignKey:identity; references:user_identity; json:"user_basic"`
	Path            string        `gorm:"column:path; type:varchar(255); json:"path"` // path of the submit file
	Status          int           `gorm:"column status; type:int; json:"status"`      // -1-not judged, 1-accepted, 2-wrong answer, 3-time limit exceeded, 4-memory limit exceeded
}

func (table *SubmitBasic) TableName() string {
	return "submit_basic"
}

func GetSubmitList(problemIdentity, userIdentity string, status int) *gorm.DB {
	tx := DB.Model(new(SubmitBasic)).Preload("ProblemBasic", func(db *gorm.DB) *gorm.DB {
		return db.Omit("content")
	}).Preload("UserBasic", func(db gorm.DB) *gorm.DB {
		return db.Omit("password")
	})
	if problemIdentity != "" {
		tx = tx.Where("problem_identity = ?", problemIdentity)
	}
	if userIdentity != "" {
		tx = tx.Where("user_identity = ?", userIdentity)
	}
	if status != 0 {
		tx = tx.Where("status = ?", status)
	}
	return tx
}
