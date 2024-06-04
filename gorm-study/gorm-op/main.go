package main

import (
	"fmt"
	"gorm-study/global"
	"gorm-study/gorm-op/model"
	"gorm-study/utils"
	"gorm.io/gorm"
	"time"
)

var db *gorm.DB

func init() {
	utils.Init()
	db = global.DB
}

func main() {
	s := model.Song{}
	//db.Find(&s) // SELECT * FROM `song` WHERE `song`.`deleted_at` IS NULL
	//db.First(&s) // SELECT * FROM `song` WHERE `song`.`deleted_at` IS NULL ORDER BY `song`.`id` LIMIT 1
	//db.Take(&s) //  SELECT * FROM `song` WHERE `song`.`deleted_at` IS NULL LIMIT 1
	//db.First(&s, 10) // SELECT * FROM `song` WHERE `song`.`id` = 10 AND `song`.`deleted_at` IS NULL ORDER BY `song`.`id` LIMIT 1

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>where

	db := db.Where("listen_count = ?", 1)

	// tody := time.Now().Format("2006-01-02 15:04:05")
	db = db.Where("created_at <= ?", time.Now())

	d, _ := time.ParseDuration("-24h")
	lastday := time.Now().Add(d)
	db = db.Where("deleted_at between ? and ?", lastday, time.Now())

	// db.Where("song_name = ? or song_id = 1", "ww").Find(&s) //((song_name = 'ww' or song_id = 1))
	db = db.Where("listen_count in (?)", []string{"1", "2"}).
		Or("listen_count = 13")

	db = db.Where("song_name like (?) AND listen_count >= ?", "%ww%", "0")

	db = db.Not("listen_count", 1)
	db = db.Not("listen_count", []int{11, 12}).Order("singer_id desc").Order("song_name")
	/**
	 SELECT * FROM `song` WHERE listen_count = 1 AND created_at <= '2024-05-31 15:26:18.28' AND (deleted_at between '2024-05-30 15:26:18.28' and '2024-05-3
	1 15:26:18.28') AND listen_count in ('1','2') OR listen_count = 13 AND (song_name like ('%ww%') AND listen_count >= '0') AND `listen_count` <> 1 AND `listen_count` NOT I
	N (11,12) ORDER BY singer_id desc,song_name
	值得注意的是Or，可以看到时直接拼接Or的，需注意优先级
	*/
	//db.Find(&s)

	//songDao := model.GetInstance()
	//singerDao := model.GetSingerDaoInstance()
	//songDao.FindSongById(108055068557582336)
	//songDao.FindSongByName("fade")
	//songDao.FindSongByStructCond(&model.Song{
	//	Id:          0,
	//	SongName:    "fade",
	//	ListenCount: 0,
	//	CreatedAt:   nil,
	//	UpdatedAt:   nil,
	//	DeletedAt:   nil,
	//}) // 零值不会作为查找条件

	//songDao.FindSongByMapCond(map[string]interface{}{
	//	"listen_count": 0,
	//}) // 零值也可以作为查找条件

	//fmt.Println("插入数据")
	//songDao.Create(&model.Song{
	//	SongName: "fade",
	//})
	//singerDao.Create(&model.Singer{
	//	Name:     "jay",
	//	NickName: "jba",
	//	Age:      40,
	//})
	//singerDao.FindAllSinger()
	//songDao.FindAllSong()

	//fmt.Println("删除")
	//singerDao.DeleteById(108098349542809600)

	//singerDao.UpdateSingerPartialById(&model.Singer{
	//	Id:       108098349542809601,
	//	Name:     "shit11",
	//	NickName: "",
	//	Age:      10,
	//})
	//songDao.IncrSongListCount(108098349408591872)

	//singerDao.UpdateSingerByIdWithMap(108098349542809601, map[string]interface{}{
	//	"name": "fuck",
	//	"age":  0,
	//})
	
	model.JoinQuery()

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>链式查询,前面查询的结果是后面的查询的范围基础
	//db.Order("singer_id desc").Find(&s).Order("song_name").Find(&s) //链式查询，相当于两次查询，第一次查询是第二次查询的范围基础

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>FirstOrInit,        仅适用与于 struct 和 map 条件

	// db.FirstOrInit(&s, model.Song{SongName: "不存在"}) //会先按`song_name` = '不存在'找，找到就找到，找不到就按写的内容赋值给s
	// db.FirstOrInit(&s, map[string]interface{}{"song_name": "xxxx"})
	// db.Where(model.Song{SongName: "ddd"}).FirstOrInit(&s)

	//Attrs，查不到就会按照Attrs赋值给结构体
	// db.Where(model.Song{SongName: "ww"}).Attrs(model.Song{SongName: "xxx", ListenCount: 30}).FirstOrInit(&s)

	//Assign
	//与FirstOrInit不同，不查不查得到都会按Assign中的属性赋值给结构体
	// db.Where(model.Song{SongId: 1}).Assign(model.Song{SongName: "ww"}).FirstOrInit(&s)

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>FirstOrCreate	  仅适用与于 struct 和 map 条件
	//与FirstOrInit类似，但是FirstOrCreate会插入数据库，找到了就更新到数据库，找不到就查到数据库，其他和FirstOrInit一样

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>子查询
	// db.Where("listen_count > (?)", db.Table("songs").Select("AVG(listen_count)").Where("singer_id = ?", "0").QueryExpr()).Find(&s)

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>Limit
	// 用 -1 取消 LIMIT 限制条件
	//var ss []model.Song
	////db.Limit(10).Find(&s).Limit(-1).Find(&ss) //链式，相当于查了两次，不加-1，第二次也会有limit限制
	//db.Offset(0).Limit(3).Find(&ss) //limit限制查询出的记录数，offset设置查找开始的下标（0开始）
	//for _, song := range ss {
	//	fmt.Println(song)
	//}

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>  Count count要放在语句最后,直接count，不需要find,但是不用find，就需要指定表
	// var count int
	// db.Table("songs").Where("song_name like (?)", "%ww%").Count(&count)
	// fmt.Println(count)

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>		group by  having  row   scan
	// rows, _ := db.Table("songs").Select("singer_id , sum(song_id) as total").Group("singer_id").Having("singer_id = 1").Rows()
	// var total int
	// defer rows.Close()
	// for rows.Next() {
	// 	rows.Scan(&s.SingerId, &total) //rows和scan一起用
	// 	fmt.Printf("singerid:%v	total:%v\n", s.SingerId, total)
	// }

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>	连接查询
	// rows, _ := db.Table("songs").Select("song_id,song_name,si.name as singer_name").Joins("left join singers si on si.mainid = songs.singer_id").Rows()
	// var songid int
	// var songname, singername string
	// for rows.Next() {
	// 	rows.Scan(&songid, &songname, &singername)
	// 	fmt.Printf("%v,%v,%v\n", songid, songname, singername)
	// }

	/*
		注意所有使用链式查询的地方，一个链中前面的查询是后面查询的基础
	*/
	fmt.Printf("%#v\n", s)
}
