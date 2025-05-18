package model

import "fmt"

// 具体音乐

type Rnb struct {
    Music // 使用匿名成员达到类似继承的效果
}

func (m *Rnb) Play() {
    fmt.Println("play -- ", m.Name, "  type:", m.Type)
}
func (m *Rnb) Stop() {
    fmt.Println("stop -- ", m.Name, "  type:", m.Type)
}

// NewRnb 这里可以其实结合单例来使用
func NewRnb() *Rnb {
    return &Rnb{Music{Name: "流沙", Type: "rnb"}}
}
