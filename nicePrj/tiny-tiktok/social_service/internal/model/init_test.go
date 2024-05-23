package model

import (
	"fmt"
	"github.com/spf13/viper"
	"testing"
	"tiny-tiktok/social_service/config"
)

func TestInitDb(t *testing.T) {
	config.InitConfig()
	dns := config.DbDnsInit()
	fmt.Print(dns)
	config.InitConfig()
	fmt.Printf("host is: %v \n", viper.GetString("mysql.host"))
}
