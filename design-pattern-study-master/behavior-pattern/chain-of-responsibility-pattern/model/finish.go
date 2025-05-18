package model

import "fmt"

// Finish 收尾中间件
type Finish struct {
}

func (c *Finish) Handle(client *Client) {
    fmt.Println("做收尾工作...")
}

func (c *Finish) SetNext(handler Handler) {
}
