package model

import "fmt"

// 具体实现类
// 具体实现类的创建可以由抽象工厂模式来实现

type Oracle struct {
    D Dsn
}

func (o *Oracle) Connect() {
    o.D.Desc()
    fmt.Println("connect to oracle successfully")
}
func (o *Oracle) SetD(d Dsn) {
    o.D = d
}
