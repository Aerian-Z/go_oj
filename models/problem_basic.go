package models

import (
	"gorm.io/gorm"
)

type ProblemBasic struct {
	gorm.Model
	Identity        string             `gorm:"column:identity; type:varchar(36); json:"identity"` // unique identification of the problem table
	Title           string             `gorm:"column:title; type:varchar(255); json:"title"`      // title of the article
	Content         string             `gorm:"column:content; type:longtext; json:"content"`      // content of the article
	MaxRuntime      int                `gorm:"column:max_runtime; type:int; json:"max_runtime"`   // max runtime of the problem
	MaxMemory       int                `gorm:"column:max_memory; type:int; json:"max_memory"`     // max memory of the problem
	ProblemCategory []*ProblemCategory `gorm:"foreignKey:problem_id; references:id; json:"problem_categories"`
}

func (table *ProblemBasic) TableName() string {
	return "problem_basic"
}

func GetProblemList(keyword, categoryIdentity string) *gorm.DB {
	//data := make([]*Problem, 0)
	//DB.Find(&data)
	//for _, v := range data {
	//	fmt.Printf("Problem ===> %v \n", v)
	//}
	tx := DB.Model(new(ProblemBasic)).Preload("ProblemCategory.CategoryBasic").
		Where("title LIKE ? OR content LIKE ? ", "%"+keyword+"%", "%"+keyword+"%")
	if categoryIdentity != "" {
		tx = tx.Joins("RIGHT JOIN problem_category pc ON pc.problem_id = problem_basic.id").
			Where("pc.category_id = (SELECT cb.id FROM category_basic cb WHERE cb.identity = ?)", categoryIdentity)
	}
	return tx
}

func GetProblemDetail(identity string) *gorm.DB {
	return DB.Model(new(ProblemBasic)).Preload("ProblemCategory.CategoryBasic").Where("identity = ?", identity)
}
