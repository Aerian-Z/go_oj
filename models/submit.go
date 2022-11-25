package models

import "gorm.io/gorm"

type Submit struct {
	gorm.Model
	Identidy        string `gorm:"column:identify; type:varchar(36); json:"identity"`                 // unique identification of the submit table
	ProblemIdentity string `gorm:"column:problem_identity; type:varchar(36); json:"problem_identity"` // unique identification of the problem table
	UserIdentity    string `gorm:"column:user_identity; type:varchar(36); json:"user_identity"`       // unique identification of the user table
	Path            string `gorm:"column:path; type:varchar(255); json:"path"`                        // path of the submit file
}

func (table *Submit) TableName() string {
	return "submit"
}
