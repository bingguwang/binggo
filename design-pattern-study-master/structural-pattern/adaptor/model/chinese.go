package model

import "fmt"

// ChinesePort 中式标准插座
type ChinesePort struct {
}

// WirelessPlug 无线充电
func (*ChinesePort) WirelessPlug() {
    fmt.Println("使用中国插座，进行无线充电，充电成功")
}
