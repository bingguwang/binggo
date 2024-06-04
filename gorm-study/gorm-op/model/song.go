package model

import (
	"encoding/json"
	"fmt"
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

	Id          int64      `gorm:"primary_key"`
	SongName    string     `gorm:"type:varchar(30)"`
	ListenCount uint       `gorm:"type:int unsigned;default:0"`
	IsHot       bool       `gorm:"type:TINYINT;default:0"` // 是否是热歌
	CreatedAt   *time.Time // 默认是0，而且插入时候会自动添加，字段名字得严格叫这个
	UpdatedAt   *time.Time // 默认是0，而且插入时候会自动添加，字段名字得严格叫这个
	// 这里最好是设为指针，查询结果映射到这里的时候，null就是null，不会是零值
	DeletedAt *time.Time `gorm:"deleted_at;index;default:null"` // 默认是0，而且插入时候会自动添加，字段名字得严格叫这个

	SingerId int64 `gorm:"BIGINT(20);default:null"`
	// song belongs to singer,假设一首歌只属于一个歌手
	// foreignkey:自己表中的关联字段;references:对方表中的关联字段
	// 然后在查询song的时候，会神奇的发现singer的值也被找了出来
	// 其实这种操作是修改了表结构，给song表新增了一个外键，指向的是singer的id
	Singer *Singer `json:"singer" gorm:"foreignkey:SingerId;references:Id"`

	// many2many 会创建一个连接表
	SongSheet []*Sheet `gorm:"many2many:sheet_songs;"` // 歌单, 一首歌属于多个歌单,每个歌单有多少歌，于是song和songsheet是多对多的关系
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

// song的dao操作

// 插入

func (*SongDao) Create(song *Song) error {
	// 雪花算法生成主键
	flake, _ := snowFlake.NewSnowFlake(7, 2) // 服务2
	song.Id = flake.NextId()

	// 单条SQL，不需要事务
	global.DB.Create(&song)

	return nil

}

// 查询
func (*SongDao) FindSongById(id uint64) *Song {
	res := &Song{}
	// 预加载，其实就是执行了两次查询
	// SELECT * FROM `singer` WHERE `singer`.`id` = 108075893650235392
	// SELECT * FROM `song` WHERE `song`.`id` = 108055068557582336 ORDER BY `song`.`id` LIMIT 1
	global.DB.Preload("Singer").First(&res, id)
	//global.DB.First(&res, id)
	fmt.Println(ToString(res))
	return res
}

func (*SongDao) FindSongByName(name string) []*Song {
	res := make([]*Song, 0)
	global.DB.Where("song_name like (?) ", "%"+name+"%").Find(&res)
	fmt.Println(len(res))
	return res
}

func (*SongDao) FindAllSong() []*Song {
	res := make([]*Song, 0)
	global.DB.Preload("Singer").Find(&res)
	//global.DB.Find(&res)
	fmt.Println(ToString(res))
	fmt.Println(len(res))
	return res
}

// FindSongByStructCond 零值不会作为查找条件
func (*SongDao) FindSongByStructCond(songCond *Song) []*Song {
	res := make([]*Song, 0)
	global.DB.Where(songCond).Find(&res)
	//global.DB.Find(&res, songCond) // 和上面等价
	fmt.Println(len(res))
	return res
}

// FindSongByMapCond 零值也可以作为查找条件
func (*SongDao) FindSongByMapCond(mapCond map[string]interface{}) []*Song {
	res := make([]*Song, 0)
	global.DB.Where(mapCond).Find(&res)
	//global.DB.Find(&res, mapCond) // 和上面等价
	fmt.Println(len(res))
	return res
}

// 更新

func UpdateQueryAndUpdateSong(id uint64, newSong *Song) error {
	db := global.DB
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	song := Song{}

	// 先读再更新的场景
	// 锁住指定 id 的 User 记录, 加for update排他锁
	if err := tx.Set("gorm:query_option", "FOR UPDATE").First(&song, id).Error; err != nil {
		// 一旦出错就回滚
		tx.Rollback()
		return err
	}

	// 更新操作...
	if song.ListenCount > 100 { // 更新为热歌
	}

	// commit事务，释放锁
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}

// IncrSongListCount 字段自增
func (*SongDao) IncrSongListCount(id uint64) error {
	db := global.DB
	db.Model(&Song{}).Where("id = ?", id).Update("listen_count", gorm.Expr("listen_count + 1"))
	return nil
}

func ToString(i interface{}) string {
	res, _ := json.Marshal(i)
	return string(res)
}
