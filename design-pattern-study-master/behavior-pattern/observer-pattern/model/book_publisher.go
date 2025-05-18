package model

import "sync"

type BookPublisher struct {
    Msg       string
    Observers []Observer
}

func (b *BookPublisher) NotifyAll() {
    var wg sync.WaitGroup
    for _, v := range b.Observers {
        wg.Add(1)
        go func(o Observer) {
            defer wg.Done()
            o.Receive(b.Msg) // 异步通知
        }(v)
    }
    wg.Wait()
}

func NewBookPublisher() Publisher {
    return &BookPublisher{Observers: []Observer{}}
}

func (b *BookPublisher) AddObserver(observer ...Observer) {
    b.Observers = append(b.Observers, observer...)
}

func (b *BookPublisher) UpdateMsg(msg string) {
    b.Msg = msg
}
