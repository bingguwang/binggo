package main

import (
	"github.com/spf13/viper"
	"net/http"
	"time"
	config "tiny-tiktok/api_router/configs"
	"tiny-tiktok/api_router/discovery"
	"tiny-tiktok/api_router/pkg/logger"
	"tiny-tiktok/api_router/router"
)

func main() {
	config.InitConfig()
	resolver := discovery.Resolver()
	r := router.InitRouter(resolver)
	server := &http.Server{
		Addr:           viper.GetString("server.port"),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		logger.Log.Fatal("启动失败...")
	}
}
