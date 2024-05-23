package main

import (
	"tiny-tiktok/video_service/config"
	"tiny-tiktok/video_service/discovery"
	"tiny-tiktok/video_service/internal/handler"
	"tiny-tiktok/video_service/internal/model"
	"tiny-tiktok/video_service/pkg/cache"
)

func main() {
	config.InitConfig() // 初始话配置文件
	model.InitDb()      // 初始化数据库
	cache.InitRedis()   // 初始化缓
	go func() {
		handler.PublishVideo()
	}()
	discovery.AutoRegister() // 自动注册
}
