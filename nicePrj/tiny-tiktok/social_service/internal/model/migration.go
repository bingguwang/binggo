// Package model 数据库自动迁移
package model

import "tiny-tiktok/social_service/pkg/logger"

func migration() {
	// 自动迁移
	err := DB.Set("gorm:table_options", "charset=utf8mb4").AutoMigrate(&Follow{}, Message{})
	// Todo 判断error 写入日志
	if err != nil {
		logger.Log.Error("err:", err.Error())
	}
}
