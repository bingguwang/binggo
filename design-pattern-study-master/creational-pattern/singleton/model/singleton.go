package model

import "sync"

// Instance 单例模式，采用懒加载，使用sync.once实现
type Instance struct{}
var (
    instance *Instance // 单例一般设为私有，且一般设为指针类型
    once     sync.Once //once中有个互斥锁，可以满足并发安全
)
// GetSingletonInstance 获取单例
func GetSingletonInstance() *Instance {
    once.Do(func() { //once只会执行一次，所以能达到单例的目的
        instance = &Instance{}
    })
    return instance
}
