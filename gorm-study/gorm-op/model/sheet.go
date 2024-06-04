package model

// 歌单，和song是多对多的关系
// 一首歌属于多个歌单
type Sheet struct {
	Id        string `gorm:"primary_key"`
	SheetName string `gorm:"type:varchar(30)"`

	//使用 AutoMigrate 为 User 创建表时，GORM 会自动创建连接表sheet_songs
	SongList []Song `gorm:"many2many:sheet_songs;references:Id"`
}
