package model

import "fmt"

// Cashier 收银中间件
type Cashier struct {
    Next Handler
}

func (c *Cashier) Handle(client *Client) {
    if client.paymentDone { // 保证只执行一次
        c.Next.Handle(client)
        return
    }
    fmt.Println("Cashier handling...")
    client.paymentDone = true
    c.Next.Handle(client)
}

func (c *Cashier) SetNext(handler Handler) {
    c.Next = handler
}
