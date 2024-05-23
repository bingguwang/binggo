package main

import (
	"tiny-tiktok/social_service/config"
	"tiny-tiktok/social_service/discovery"
	"tiny-tiktok/social_service/internal/model"
	"tiny-tiktok/social_service/pkg/cache"
)

func main() {
	config.InitConfig()
	model.InitDb()
	cache.InitRedis()
	go cache.TimerSync()
	discovery.AutoRegister()
}
