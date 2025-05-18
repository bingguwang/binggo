package model

import "fmt"

// 音乐产品，加这一层是为了模拟继承的效果

type Music struct {
    Name string
    Type string
}

func (m *Music) Play() {
    fmt.Println("play -- ", m.Name)
}
func (m *Music) Stop() {
    fmt.Println("stop -- ", m.Name)
}
