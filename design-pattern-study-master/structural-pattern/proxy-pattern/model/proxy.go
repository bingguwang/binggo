package model

import "fmt"

type Proxy struct {
    User *User
}

// 貌似装饰器模式和静态代理没啥区别啊

func (p *Proxy) Login() {
    fmt.Println("login by proxy pre operation")
    p.User.Login()
    fmt.Println("login by proxy after operation")
}

func NewProxy(user *User) *Proxy {
    return &Proxy{User: user}
}
