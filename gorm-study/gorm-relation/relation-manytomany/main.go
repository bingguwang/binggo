package main

//many 2 many //多对多，一般处理多对多是建立一个连接表
import (
	"fmt"
	"gorm-study/global"
	"gorm-study/gorm-relation/relation-manytomany/model"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var db *gorm.DB

func main() {
	db = global.DB

	var singers []model.Singer
	fmt.Println(singers)

	// db.AutoMigrate(&model.Sheet{})

}
