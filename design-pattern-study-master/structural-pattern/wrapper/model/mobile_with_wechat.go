package model

import "fmt"

// 装饰器2

type MobileWithWeChat struct {
    Mobile IMobile
}

func (m *MobileWithWeChat) Function() {
    m.Mobile.Function()
    fmt.Println("play wechat")
}
