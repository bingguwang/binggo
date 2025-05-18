package model

import "fmt"

type OrderItem struct {
    Name string
}

func (o *OrderItem) Desc() {
    fmt.Println("this is a order item ：", o.Name)
}
