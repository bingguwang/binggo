package main

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm-study/global"
	"gorm-study/gorm-model/model"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	db = global.DB
	// 自动迁移,使用gorm自动对应实体创建表结构，
	// 仅支持创建表、增加表中没有的字段和索引。为了保护你的数据，它并不支持改变已有的字段类型或删除未被使用的字段
	db.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&model.Song{})

	s := model.Song{ListenCount: 0} //如果是存的0值的话，是不会被插入的,而是取默认值
	db.Create(&s)

}
