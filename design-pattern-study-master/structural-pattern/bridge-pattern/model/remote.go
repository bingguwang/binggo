package model

import "fmt"

type Remote struct {
}

func (r *Remote) Desc() {
    fmt.Println("remote")
}
