package database

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var dbInstance *gorm.DB

func InitDB() error {

	source := "root:buqieryu@tcp(mysql:3306)/buqieryu?readTimeout=1500ms&writeTimeout=1500ms&charset=utf8&loc=Local&&parseTime=true"
	// user := os.Getenv("MYSQL_USERNAME")
	// pwd := os.Getenv("MYSQL_PASSWORD")
	// addr := os.Getenv("MYSQL_ADDRESS")
	// dataBase := os.Getenv("MYSQL_DATABASE")
	// if dataBase == "" {
	// 	dataBase = "golang_demo"
	// }
	// source = fmt.Sprintf(source, user, pwd, addr, dataBase)
	// fmt.Println("start init mysql with ", source)

	db, err := gorm.Open(mysql.Open(source), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		}})
	if err != nil {
		fmt.Println("DB Open error,err=", err.Error())
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("DB Init error,err=", err.Error())
		return err
	}

	// 用于设置连接池中空闲连接的最大数量
	sqlDB.SetMaxIdleConns(100)
	// 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(200)
	// 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	dbInstance = db

	fmt.Println("finish init mysql with ", source)
	return nil
}

// func InitLocalDB() error {
// 	source := "root:wwb20030526@tcp(127.0.0.1:3308)/test?readTimeout=1500ms&writeTimeout=1500ms&charset=utf8&loc=Local&&parseTime=true"
// 	db, err := gorm.Open(mysql.Open(source), &gorm.Config{
// 		NamingStrategy: schema.NamingStrategy{
// 			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
// 		}})
// 	if err != nil {
// 		fmt.Println("DB Open error,err=", err.Error())
// 		return err
// 	}

// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		fmt.Println("DB Init error,err=", err.Error())
// 		return err
// 	}
// 	sqlDB.SetMaxIdleConns(100)
// 	// 设置打开数据库连接的最大数量
// 	sqlDB.SetMaxOpenConns(200)
// 	// 设置了连接可复用的最大时间
// 	sqlDB.SetConnMaxLifetime(time.Hour)

// 	dbInstance = db

// 	fmt.Println("finish init mysql with ", source)
// 	return nil
// }

// Get ...
func Get() *gorm.DB {
	return dbInstance
}
