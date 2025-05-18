package model

import "fmt"

type SonyRock struct {
    Music
}

func (m *SonyRock) Play() {
    fmt.Println("play -- ", m.Name, "  type:", m.Type, "  copyright:", m.Copyright)
}
func (m *SonyRock) Stop() {
    fmt.Println("stop -- ", m.Name, "  type:", m.Type, "  copyright:", m.Copyright)
}

// NewSonyRock 这里可以其实结合单例来使用
func NewSonyRock() *SonyRock {
    return &SonyRock{Music{Name: "黑色柳丁", Type: "rock", Copyright: "sony"}}
}
