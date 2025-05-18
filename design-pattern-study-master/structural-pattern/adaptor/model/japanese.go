package model

import "fmt"

type JapanesePort struct {
}

func (*JapanesePort) CommonPlug() {
    fmt.Println("使用日本插座普通充电，充电成功")
}
