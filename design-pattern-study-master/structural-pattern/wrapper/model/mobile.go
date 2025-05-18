package model

import "fmt"

type Mobile struct {
}

func (*Mobile) Function() {
    fmt.Println("call")
}
