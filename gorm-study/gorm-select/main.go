package main

import (
	"fmt"
	"gorm-study/global"
	"gorm-study/gorm-select/model"
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

	//db.Where("singer_id = ?", 1).First(&s)
	// db.Where("singer_id = ?", 1).Find(&s)
	/**
	SELECT * FROM `song` WHERE song_id in ('1','2') AND `song`.`d	deleted_at` IS NULL
	*/
	//db.Where("song_id in (?)", []string{"1", "2"}).Find(&s)

	// db.Where("song_name like ?", "%ww%").Find(&s)
	// db.Where("song_name like (?) AND listen_count >= ?", "%ww%", "0").Find(&s)

	// tody := time.Now().Format("2006-01-02 15:04:05")
	db.Where("created_at <= ?", time.Now()).Find(&s)

	// d, _ := time.ParseDuration("-24h")
	// lastday := tody.Add(d)
	// db.Where("created_at between ? and ?", lastday, tody).Find(&s)

	// //where可以传结构体
	// db.Where(&model.Song{SongName: "ww", SongId: 1}).Find(&s) //当含有传入的结构体体含零值的时候，含零值的条件会被忽略

	// //where可以传map
	// db.Where(map[string]interface{}{"song_name": "ww", "song_id": 2}).Find(&s)

	//Not，和 Where查询类似
	// db.Not("song_id", 1).Find(&s)                //条件相当于`song_id` NOT IN (1)
	// db.Not("song_id", []int{1, 2}).Find(&s)      //条件相当于`song_id` NOT IN (1,2)
	// db.Not(&model.Song{SongName: "ww"}).Find(&s) //song_name` <> 'ww'

	// db.Where("song_name = ? or song_id = 1", "ww").Find(&s) //((song_name = 'ww' or song_id = 1))
	// db.Where("song_name = ?", "ww").Or("song_id = 1").Find(&s) //((song_name = 'ww') OR (song_id = 1))
	// db.Where(&model.Song{SongName: "ww"}).Or(&model.Song{SongId: 1}).Find(&s)

	// db.Find(&s, &model.Song{SongName: "ww", SingerId: 1})
	// db.Find(&s, map[string]interface{}{"Song_name": "ww", "Song_id": 1})

	// db.Find(&s, 1) //where 主键=1

	// 为查询 SQL 添加额外的选项
	// db.Set("gorm:query_option", "FOR UPDATE").First(&s, 1) //select ....for update

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>排序order
	// db.Order("singer_id desc").Order("song_name").Find(&s) //ORDER BY singer_id desc,song_name
	// db.Order("singer_id desc").Find(&s).Order("song_name").Find(&s)       //链式查询，相当于两次查询，第一次查询是第二次查询的范围基础
	// db.Order("singer_id desc").Find(&s).Order("song_name", true).Find(&s) /*和上一句不同是第二次查询的排序不同
	// 上面是ORDER BY singer_id desc,song，下面加了true的是 ORDER BY song_name,
	// */

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
	var ss []model.Song
	//db.Limit(10).Find(&s).Limit(-1).Find(&ss) //链式，相当于查了两次，不加-1，第二次也会有limit限制
	db.Offset(0).Limit(3).Find(&ss) //limit限制查询出的记录数，offset设置查找开始的下标（0开始）
	for _, song := range ss {
		fmt.Println(song)
	}

	//>>>>>>>>>>>>>>>>>>>>>>>>>>>>	查询指定字段 	select

	//  SELECT song_name, song_id FROM `song` WHERE song_name like ('%ww%') AND `song`.`deleted_at` IS NULL
	db.Select("song_name, song_id").Where("song_name like (?)", "%ww%").Find(&ss)

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
