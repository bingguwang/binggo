package model

import "fmt"

type Local struct {
}

func (r *Local) Desc() {
    fmt.Println("local")
}
