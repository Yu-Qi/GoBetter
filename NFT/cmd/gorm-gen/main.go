package main

import (
	"github.com/Yu-Qi/GoBetter/NFT/pkg/db/model"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	db, _ := gorm.Open(mysql.Open("root:changeit@(127.0.0.1:3306)/GoBetter?charset=utf8mb4&parseTime=True&loc=Local"))
	g := gen.NewGenerator(gen.Config{
		OutPath: "pkg/db/dao",
	})
	g.UseDB(db)
	g.GenerateAllTable()

	g.ApplyBasic(model.User{})
	g.Execute()

}
