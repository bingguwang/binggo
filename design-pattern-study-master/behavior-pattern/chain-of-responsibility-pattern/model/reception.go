package model

import "fmt"

// Reception 中间件
type Reception struct {
    Next Handler
}

func (c *Reception) Handle(client *Client) {
    if client.registrationDone { // 保证只执行一次
        c.Next.Handle(client)
        return
    }
    fmt.Println("Reception handling...")
    client.registrationDone = true
    c.Next.Handle(client)
}

func (c *Reception) SetNext(handler Handler) {
    c.Next = handler
}
