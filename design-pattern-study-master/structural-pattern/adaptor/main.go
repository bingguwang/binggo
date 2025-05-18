package main

import (
    "fmt"
    "github.com/BingguWang/design-pattern-study/structural-pattern/adaptor/model"
)

/**
如果我希望使用某个类， 但是其接口与其他代码不兼容时，怎么办？

适配器模式
    适配器是为了转换对象，使得对象可以满足与其他对象交互

适配器模式通过
    封装对象将复杂的转换过程隐藏于幕后。 被封装的对象甚至察觉不到适配器的存在。

适配器有类适配器和对象适配器
类适配器：适配器同时继承两个对象的接口。 这种方式仅能在支持多重继承的编程语言中实现

go没有继承的概念，所以它实现的适配器一般都是对象适配器
下面演示的是对象适配器
*/
func main() {
    client := model.Client{}
    fmt.Println("客户调用日式插座充电")
    client.Charge(&model.JapanesePort{})

    fmt.Println("客户调用标准插座")
    client.Charge(&model.Adaptor{})

    /**
      client只知道调用标准插座充电，对于中式插座的存在它根本不知道，它所知道的是 只要传入适配器它就可以调到无线充电
    */
}
