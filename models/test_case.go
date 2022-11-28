package models

import "gorm.io/gorm"

type TestCase struct {
	gorm.Model
	Identity        string `gorm:"column:identity; type:varchar(36); json:"identity"`                 // unique identification of the test case table
	ProblemIdentity string `gorm:"column:problem_identity; type:varchar(36); json:"problem_identity"` // unique identification of the problem table
	Input           string `gorm:"column:input; type:varchar(255); json:"input"`
	Output          string `gorm:"column:output; type:varchar(255); json:"output"`
}

func (table *TestCase) TableName() string {
	return "test_case"
}
