package test

import (
	"fmt"
	"go_oj/models"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestGormTest(t *testing.T) {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:@tcp(127.0.0.1:3306)/go-oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// fmt.Println(db, err)
	if err != nil {
		t.Fatal(err)
	}

	data := make([]*models.ProblemBasic, 0)
	err = db.Find(&data).Error
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range data {
		fmt.Printf("Problem ===> %v \n", v)
	}
}
