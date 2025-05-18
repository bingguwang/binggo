package model

type Handler interface {
    Handle(*Client)
    SetNext(handler Handler)
}
