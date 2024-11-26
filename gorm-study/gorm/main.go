package main

import (
	"fmt"
	"gorm-study/gorm/gormutils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// orm的作用主要是映射程序字段和数据库字段，方便持久化
// 数据库驱动依赖不能少
func main() {
	//拼接下dsn参数, dsn格式可以参考上面的语法，这里使用Sprintf动态拼接dsn参数，因为一般数据库连接参数，我们都是保存在配置文件里面，需要从配置文件加载参数，然后拼接dsn。
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", `root`, `12345`, `127.0.0.1`, `3306`, `test`, 5)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		Logger: logger.Default,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 默认不加负数
		},
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。

	// db.SetLogger(logDemo.New(os.Stdout, "\r\n", 0)) //设置日志格式

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>插入操作>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	s := gormutils.Singer{Name: "周杰伦", NickName: "周董"} //满足驼峰命名方法
	//插入数据
	db.Create(&s)

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>删除操作>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	//删除数据
	var ss gormutils.Singer
	db.Where("name = ?", "谭咏麟").Delete(&ss)

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>更新操作>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	//修改数据
	db.Model(&ss).Where("mainid = ?", 8).Update("name", "xxx") //每个Update就是一个update语句
	// UPDATE `singers` SET `name` = 'xxx'  WHERE (mainid = 8)

	//查出来再修改，可以直接用查询语句和save语句
	sg := gormutils.Singer{}
	db.Where("mainid = ?", 3).Take(&sg)
	sg.SetName("tiger1")
	sg.SetNickName("虎子1")
	db.Model(&sg).Where("mainid = ?", sg.MainId).Update(&sg) //where和update方法顺序不能反了

	//updates跟新多个字段
	mp := make(map[string]interface{})
	mp["name"] = "小青"
	mp["nickName"] = "妖怪"
	db.Model(&sg).Where("mainid > 10").Updates(mp)

	//UPDATE singer SET mainid = mainid + 1 WHERE id = '20' //SQL中有表达式的
	db.Model(&sg).Where("mainid = 20").Update("age", gorm.Expr("age + 1"))

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>>事务>>>>>>>>>>>>>>>>>>>>>>>>>>>>>")
	//类似database/sql包的事务，
	//开启事务
	tx := db.Begin()
	ad := gormutils.Singer{
		MainId:   1, //此主键已存在，会报错，看下是否能回滚
		Name:     "陶喆",
		Age:      51,
		NickName: "DT",
	}
	if err := tx.Create(&ad); err != nil {
		//回滚
		tx.Rollback() //回滚后后面的所有SQL都会失效，包括查询也查不出
		//如果不在这调用rollback，那就是错误的SQL不会执行，但是其他正确的操作会执行，那就没有回滚的效果了，
		// 所以还是要执行rollback才对
		fmt.Println(err)
	}

	var bd gormutils.Singer
	tx.Where("mainid = ?", 2).Take(&bd) //即使是查询操作还是不会操作，因为前面有回滚
	fmt.Println(bd)
	bd.SetName("刘德华")
	tx.Model(&bd).Where("mainid = ?", bd.GetMainId()).Update(&bd)

	//提交事务
	tx.Commit()

}
