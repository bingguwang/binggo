package model

// 具体音乐

type Rock struct {
    Music
}

func NewRock() *Rock {
    return &Rock{Music{Name: "黑色柳丁", Type: "rock"}}
}
