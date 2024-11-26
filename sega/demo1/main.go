package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dtm-labs/dtmcli"
	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid/v3"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var (
	db          *sql.DB
	redisClient *redis.Client
)

func init() {
	// 先创建一个数据库表
	var err error
	username := "root"   //数据库用户名
	password := "123456" //数据库密码
	host := "localhost"  //数据库主机号
	part := 3306         //数据库端口号
	Dbname := "test"     //数据库名称
	//root:root@tcp(127.0.0.1:3306)/gorm?
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, part, Dbname)
	//连接MYSQL,获得DB类型实例，用于后面的数据库读写操作
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("数据库连接失败，err=" + err.Error())
	}
	//连接成功
	fmt.Println(db)
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(2) //设置连接池，空闲
	sqlDB.SetMaxOpenConns(5) //设置打开最大连接

	// 自动迁移
	if err := db.AutoMigrate(&UserAccount{}); err != nil {
		panic("自动迁移失败:" + err.Error())
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
}

type UserAccount struct {
	Id         int64     `gorm:"id"`
	UserId     int       `gorm:"user_id"`
	Balance    float64   `gorm:"balance"`
	CreateTime time.Time `gorm:"create_time"`
	UpdateTime time.Time `gorm:"update_time"`
}

const dtmServer = "http://localhost:36789/api/dtmsvr" // DTM服务运行的端口
const qsBusiAPI = "/api/busi_start"
const qsBusiPort = 8082

var qsBusi = fmt.Sprintf("http://localhost:%d%s", qsBusiPort, qsBusiAPI)

func main() {
	QsStartSvr()
	QsFireRequest()
	select {}
}

// QsStartSvr quick start: start server
func QsStartSvr() {
	app := gin.New()
	qsAddRoute(app)
	log.Printf("quick start examples listening at %d", qsBusiPort)
	go func() {
		_ = app.Run(fmt.Sprintf(":%d", qsBusiPort))
	}()
	time.Sleep(100 * time.Millisecond)
}

func qsAddRoute(app *gin.Engine) {
	app.POST(qsBusiAPI+"/TransIn", func(c *gin.Context) {
		log.Printf("TransIn")
		c.JSON(200, "")
		// c.JSON(409, "") // Status 409 for Failure. Won't be retried
	})
	app.POST(qsBusiAPI+"/TransInCompensate", func(c *gin.Context) {
		log.Printf("TransInCompensate")
		c.JSON(200, "")
	})
	app.POST(qsBusiAPI+"/TransOut", func(c *gin.Context) {
		log.Printf("TransOut")
		c.JSON(200, "")
	})
	app.POST(qsBusiAPI+"/TransOutCompensate", func(c *gin.Context) {
		log.Printf("TransOutCompensate")
		c.JSON(200, "")
	})
}

// QsFireRequest quick start: fire request
func QsFireRequest() string {
	req := &gin.H{"amount": 30} // load of micro-service
	saga := dtmcli.NewSaga(dtmServer, shortuuid.New()).
		// 添加一个TransOut的子事务，正向操作为url: qsBusi+"/TransOut"， 补偿操作为url: qsBusi+"/TransOutCompensate"
		Add(qsBusi+"/TransOut", qsBusi+"/TransOutCompensate", req).
		// 添加一个TransIn的子事务，正向操作为url: qsBusi+"/TransIn"， 补偿操作为url: qsBusi+"/TransInCompensate"
		Add(qsBusi+"/TransIn", qsBusi+"/TransInCompensate", req)
	// 提交saga事务，dtm会完成所有的子事务/回滚所有的子事务
	if err := saga.Submit(); err != nil {
		panic(err)
	}
	return saga.Gid
}

func RedisOp() error {
	err := redisClient.Set(context.Background(), "test", "1", 0).Err()
	return err
}
func DBOp() error {
	return nil
}
