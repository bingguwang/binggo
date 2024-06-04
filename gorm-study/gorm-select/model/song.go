package model

import (
	"gorm-study/global"
	"gorm-study/utils/snowFlake"
	"gorm.io/gorm"
	"sync"
	"time"
)

type Song struct {
	//gorm.Model // 将字段 `ID`, `CreatedAt`, `UpdatedAt`, `DeletedAt` 注入到 `User` 模型中
	//创建的时候会默认把Model中的ID作为主键，想用自己的字段主键加上;"primary_key"就可以了
	//SongId      int64  `gorm:"type:int;unique_index;not null;primary_key"`
	//自增ID是有缺点的，对分布式不试用，所以一般使用雪花id

	Id          int64  `gorm:primary_key`
	SongName    string `gorm:"type:varchar(30)"`
	SingerId    uint   `gorm:"type:int unsigned"`
	ListenCount uint   `gorm:"type:int unsigned;default:0"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

// 只要实现Tabler接口就能自定义表名
func (*Song) TableName() string {
	return "song" // 返回你要自定义的表名，不自定义默认会加s
}

type SongDao struct {
}

var songDao *SongDao
var userOnce sync.Once // 单例模式

// GetInstance 获取单例实例
func GetInstance() *SongDao {
	userOnce.Do(
		func() {
			songDao = &SongDao{}
		},
	)
	return songDao
}

func (*SongDao) Create(song *Song) error {
	// 雪花算法生成主键
	flake, _ := snowFlake.NewSnowFlake(7, 2) // 服务2
	song.Id = flake.NextId()

	global.DB.Create(&song)

	return nil
}
