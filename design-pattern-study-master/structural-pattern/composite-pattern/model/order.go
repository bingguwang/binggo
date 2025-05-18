package model

import "fmt"

type Order struct {
    Name       string
    Components []Component // 组合模式， 有层级关系
}

func (o *Order) Desc() {
    fmt.Println("this is a order: ", o.Name)
    for _, v := range o.Components {
        v.Desc()
    }
}
