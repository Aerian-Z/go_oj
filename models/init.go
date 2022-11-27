package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB = Init()

func Init() *gorm.DB {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:@tcp(127.0.0.1:3306)/go-oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// fmt.Println(db, err)
	if err != nil {
		log.Println("gorm init error: ", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
	}
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	db.AutoMigrate(&CategoryBasic{})
	db.AutoMigrate(&ProblemBasic{})
	db.AutoMigrate(&SubmitBasic{})
	db.AutoMigrate(&UserBasic{})
	db.AutoMigrate(&ProblemCategory{})
	return db
}
