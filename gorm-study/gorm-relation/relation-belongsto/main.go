package main

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm-study/global"
	"gorm-study/gorm-relation/relation-belongsto/model"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	db = global.DB
	var songs []model.Song
	db.Preload("Singer").Find(&songs) //默认是左连接的
	for _, v := range songs {
		v.ToString()
	}

}
