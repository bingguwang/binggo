package model

// Observer 观察者，也叫订阅者
type Observer interface {
    Desc()
    Receive(msg string)
}
