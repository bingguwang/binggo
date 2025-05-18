package model

type Publisher interface {
    AddObserver(observer ...Observer)
    NotifyAll()
}
