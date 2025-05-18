package main

import (
    "fmt"
    "github.com/BingguWang/design-pattern-study/creational-pattern/singleton/model"
    "sync"
)

//懒加载，用sync.Once实现并发安全的单例
func main() {
    //testSingleton()
    testSingletonParallel()
}

//单线程判断是否是单例
func testSingleton() {
    //获取单例对象
    apple1 := model.GetSingletonInstance()
    apple2 := model.GetSingletonInstance()
    fmt.Printf("是否是单例:%v\n", apple1 == apple2)

}

//多线程下判断是否还是单例
const count = 100

func testSingletonParallel() {
    var wait sync.WaitGroup
    wait.Add(count)
    appleList := make([]*model.Instance, count)
    for i := 0; i < count; i++ {
        go func(index int) {
            //fmt.Printf("%T\n",appleList)
            appleList[index] = model.GetSingletonInstance()
            wait.Done()
        }(i)
    }
    wait.Wait()
    for i := 1; i < count; i++ {
        if appleList[i-1] != appleList[i] {
            fmt.Println("不满足单例")
            return
        }
    }
    fmt.Println("满足单例")
}
