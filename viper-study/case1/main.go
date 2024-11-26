package main

import (
	"case1/config"
	"fmt"
)

func init() {
	config.Viper()
}

func main() {
	fmt.Println(config.GVA_CONFIG)
}
