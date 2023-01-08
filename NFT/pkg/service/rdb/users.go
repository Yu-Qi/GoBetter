package rdb

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var RestfulDB *gorm.DB

func init() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:changeit@tcp(127.0.0.1:3306)/GoBetter?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	RestfulDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}
