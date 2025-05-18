package model

import "fmt"

type BookObserver struct {
    Name string
}

func (b *BookObserver) Desc() {
    fmt.Println(b.Name)
}
func (b *BookObserver) Receive(msg string) {
    fmt.Println(b.Name, " receive msg : ", msg)
}
