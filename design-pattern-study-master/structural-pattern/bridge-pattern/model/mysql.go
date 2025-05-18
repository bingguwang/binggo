package model

import "fmt"

type Mysql struct {
    D Dsn
}

func (m *Mysql) Connect() {
    m.D.Desc()
    fmt.Println("connect to mysql successfully")
}
func (m *Mysql) SetD(d Dsn) {
    m.D = d
}
