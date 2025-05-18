package model

import "fmt"

type SonyRnb struct {
    Music
}

func (m *SonyRnb) Play() {
    fmt.Println("play -- ", m.Name, "  type:", m.Type, "  copyright:", m.Copyright)
}
func (m *SonyRnb) Stop() {
    fmt.Println("stop -- ", m.Name, "  type:", m.Type, "  copyright:", m.Copyright)
}

// NewSonyRnb 这里可以其实结合单例来使用
func NewSonyRnb() *SonyRnb {
    return &SonyRnb{Music{Name: "流沙", Type: "rnb", Copyright: "sony"}}
}
