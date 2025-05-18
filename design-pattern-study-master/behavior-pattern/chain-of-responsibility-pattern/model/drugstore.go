package model

import "fmt"

// Drugstore 中间件
type Drugstore struct {
    Next Handler
}

func (c *Drugstore) Handle(client *Client) {
    if client.getMedicineDone { // 此判断保证只执行一次
        c.Next.Handle(client)
        return
    }
    fmt.Println("Drugstore handling...")
    client.getMedicineDone = true
    c.Next.Handle(client)
}

func (c *Drugstore) SetNext(handler Handler) {
    c.Next = handler
}
