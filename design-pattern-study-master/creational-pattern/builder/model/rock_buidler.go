package model

import "fmt"

type RockBuilder struct {
}

func (r *RockBuilder) BuildMusicStep1(music *Music) {
    fmt.Println("创建rock 的 第一步")
    music.Name = "黑色柳丁"
}

func (r *RockBuilder) BuildMusicStep2(music *Music) {
    fmt.Println("创建rock 的 第二步")
    music.Type = "rock"
}

func NewRockBuilder() IBuilder {
    return &RockBuilder{}
}
