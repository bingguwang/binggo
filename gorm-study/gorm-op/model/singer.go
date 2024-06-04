package model

import (
	"fmt"
	"gorm-study/global"
	"gorm-study/utils/snowFlake"
	"sync"
	"time"
)

type Singer struct {
	Id        int64      `gorm:"primary_key"`
	Name      string     `gorm:"type:varchar(100)"`
	NickName  string     `gorm:"type:varchar(100)"`
	Age       int        `gorm:"type:int(10)"`
	CreatedAt *time.Time // 默认是0，而且插入时候会自动添加，字段名字得严格叫这个
	UpdatedAt *time.Time // 默认是0，而且插入时候会自动添加，字段名字得严格叫这个
	// 这里最好是设为指针，查询结果映射到这里的时候，null就是null，不会是零值
	DeletedAt *time.Time `gorm:"deleted_at;index;default:null"` // 默认是0，而且插入时候会自动添加，字段名字得严格叫这个

	// has many 一对多,foreignKey:对方表中的关联关键字;
	// 然后在查询Singer的时候，会神奇的发现每个singer所有的song也被找了出来，
	// 这种操作也是会取修改song表的表结构，而且也会创建
	// 不会修改singer的表结构，给song表新增了一个外键，指向的是singer的id！！给song表加了个一样的外键
	// 所以，其实这里has many和 belongs to值设置一个tag就行了!!!，不然会创建一个一模一样的外键，
	//SongList []*Song `gorm:"foreignKey:SingerId;"`
	SongList []*Song

	// 请注意，如果你在迁移的时候设置了不创建外键，那整个数据库都不会自动创建外键，连接表里也不会有外键
	// 而且！预加载仍然是有效的！！！！！！！！
	// ----------------- -----------------
	//注意^-^

	// ----------------- -----------------
}

// 只要实现Tabler接口就能自定义表名
func (*Singer) TableName() string {
	return "singer" // 返回你要自定义的表名，不自定义默认会加s
}

type SingerDao struct {
}

var singerDao *SingerDao
var singerOnce sync.Once // 单例模式

// GetInstance 获取单例实例
func GetSingerDaoInstance() *SingerDao {
	singerOnce.Do(
		func() {
			singerDao = &SingerDao{}
		},
	)
	return singerDao
}

// 插入

func (*SingerDao) Create(singer *Singer) error {
	// 雪花算法生成主键
	flake, _ := snowFlake.NewSnowFlake(7, 2) // 服务2
	singer.Id = flake.NextId()

	// 单条SQL，不需要事务
	global.DB.Create(&singer)

	return nil

}

// 查询

func (*SingerDao) FindAllSinger() []*Singer {
	res := make([]*Singer, 0)
	global.DB.Preload("SongList").Find(&res)
	//global.DB.Find(&res)
	fmt.Println(ToString(res))
	fmt.Println(len(res))
	return res
}

func JoinQuery() {
	var results []struct {
		SingerName string
		SongName   string
	}
	db := global.DB
	db.Table("singer").
		Select("singer.name as singer_name, song.song_name as song_name").
		Joins("left join song on song.singer_id = song.id").
		//Where("orders.amount > ?", 100).
		Scan(&results)
	for _, result := range results {
		fmt.Println(ToString(result))
	}
	// SELECT singer.name as singer_name, song.song_name as song_name FROM `singer`
	// left join song on song.singer_id = song.id
}

// 删除

func (*SingerDao) DeleteById(id uint64) error {
	// 单条SQL，不需要事务
	global.DB.Where("id = ?", id).Delete(&Singer{})
	return nil
}

// 更新

func (*SingerDao) UpdateSinger(newSinger *Singer) error {
	db := global.DB
	db.Save(&newSinger) // 此时newSinger里的零值会被采用
	// 如果newSinger包含了主键的值，那就是update，否则就是insert操作
	// 主键是零值相当于没有设置主键，就是insert
	return nil
}

func (*SingerDao) UpdateSingerById(id int64, newSinger *Singer) error {
	db := global.DB
	newSinger.Id = id   // 一般来说，希望插入就是插入，更新就是更新，这样写就行了
	db.Save(&newSinger) // 此时newSinger里的零值会被采用
	return nil
}

func (*SingerDao) UpdateSingerByIdWithMap(id int64, newSinger map[string]interface{}) error {
	db := global.DB
	db.Model(&Singer{Id: id}).Updates(newSinger) // 可以设置零值，会被采用
	return nil
}

// 只更新想更新的字段
func (*SingerDao) UpdateSingerPartialById(newSinger *Singer) error {
	db := global.DB
	db.Model(&newSinger).Updates(&newSinger) // 只要newSinger里设置了id才会更新成功，而且零值会被忽略
	return nil
}
