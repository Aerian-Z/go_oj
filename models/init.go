package models

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB = Init()

var RDB = InitRedisDB()

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
	db.AutoMigrate(&TestCase{})
	return db
}

func InitRedisDB() *redis.Client {
	var rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}
