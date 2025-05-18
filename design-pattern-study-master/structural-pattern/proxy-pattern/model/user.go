package model

import "fmt"

type IUser interface {
    Login()
}

type User struct {
}

func (u *User) Login() {
    fmt.Println("login user")
}
