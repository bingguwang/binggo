package main

//has many  一对多
import (
	_ "github.com/go-sql-driver/mysql"
	"gorm-study/global"
	"gorm-study/gorm-relation/relation-hasmany/model"
	"gorm-study/utils"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	utils.Init()
	db = global.DB
}
func main() {
	//var singers []model.Singer
	//
	//db.Preload("SongList").Find(&singers)
	//for _, v := range singers {
	//	v.ToString()
	//	for _, s := range v.SongList {
	//		s.ToString()
	//	}
	//	fmt.Println("-----------------------")
	//}

	singer := model.Singer{
		SongList: []model.Song{
			{
				SongName: "fuckkkkkkk",
			},
			{
				SongName: "shittttttt",
			},
		},
	}
	db.Omit("nick_name").Create(&singer)

}
