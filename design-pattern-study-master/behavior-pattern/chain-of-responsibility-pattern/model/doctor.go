package model

import "fmt"

// Doctor 中间件
type Doctor struct {
    Next Handler
}

func (c *Doctor) Handle(client *Client) {
    if client.doctorCheckUpDone { // 保证只执行一次
        c.Next.Handle(client)
        return
    }
    fmt.Println("Doctor handling...")
    client.doctorCheckUpDone = true
    c.Next.Handle(client)
}

func (c *Doctor) SetNext(handler Handler) {
    c.Next = handler
}
