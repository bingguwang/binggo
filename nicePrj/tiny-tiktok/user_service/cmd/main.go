package main

import (
	"tiny-tiktok/user_service/config"
	"tiny-tiktok/user_service/discovery"
	"tiny-tiktok/user_service/internal/model"
	"tiny-tiktok/user_service/pkg/cache"
)

func main() {
	config.InitConfig()      // 初始话配置文件
	cache.InitRedis()        // 初始化redis
	model.InitDb()           // 初始化数据库
	discovery.AutoRegister() // 自动注册
}
