package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf // rest服务配置
	// todo 还有其他类型服务的配置也写在这
}
