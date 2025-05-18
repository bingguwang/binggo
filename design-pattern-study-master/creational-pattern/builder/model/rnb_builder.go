package model

import "fmt"

type RnbBuilder struct {
}

func (r *RnbBuilder) BuildMusicStep1(music *Music) {
    fmt.Println("创建rnb 的 第一步")
    music.Name = "流沙"
}

func (r *RnbBuilder) BuildMusicStep2(music *Music) {
    fmt.Println("创建rnb 的 第二步")
    music.Type = "rnb"
}

func NewRnbBuilder() IBuilder {
    return &RnbBuilder{}
}
