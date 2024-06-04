package model

// 歌单，和song是多对多的关系
type Sheet struct {
	Id        string `gorm:"primary_key"`
	SheetName string `gorm:"type:varchar(30)"`
	SongList  []Song `gorm:"many2many:sheet_songs;references:songId"` //使用 AutoMigrate 为 User 创建表时，GORM 会自动创建连接表sheet_songs
}
