package utils

import (
	"fmt"
	"gorm-study/global"
	"gorm-study/gorm-op/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"strings"
	"time"
)

const (
	ip       = "192.168.0.58"
	port     = `3306`
	username = "root"
	passwd   = "123456"
	dbname   = "test"
)

func Init() {
	dsn := strings.Join([]string{username, ":", passwd, "@tcp(", ip, ":", port, ")/", dbname, "?charset=utf8&parseTime=True&loc=Local"}, "")
	fmt.Println(dsn)

	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true, // 数据迁移时不生成外键!
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: false,
			Colorful:                  true,
		}),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 默认不加负数
		},
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	sqlDB, _ := db.DB()

	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。

	global.DB = db
	migration()
}

func migration() {
	// 自动迁移
	if err := global.DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&model.Song{}); err != nil {
		panic("自动迁移失败:" + err.Error())
	}
}
