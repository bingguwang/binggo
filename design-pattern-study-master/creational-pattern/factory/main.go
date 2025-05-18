package main

import "github.com/BingguWang/design-pattern-study/creational-pattern/factory/model"

/**
简单工厂
    使用特殊的工厂方法代替对于对象构造函数的直接调用。

    换言之，就是把"对象构造函数的调用"放到了工厂方法里而已，工厂方法返回的对象通常被称作 “产品”。

    适用场景:
        处理大型资源密集型对象 （比如数据库连接、 文件系统和网络资源） 时
        当需要有一个既能够创建新对象， 又可以重用现有对象的普通方法时，此时一般结合单例模式来实现

工厂方法做到了 把产品的创建和使用分离开，比如这里，我们要加一个jazz音乐我需要这样做：
    写一个jazz struct， 实现一下iMusic接口， 修改一下工厂方法就行了

可以看到还是有不足，就是我要去修改工厂方法，所以称为简单工厂模式，其实本质上产品之间还是耦合在一起的，耦合在这个工厂方法里
*/
func main() {
    //rockmus := model.FactoryToProduct("rnb")
    factory := model.NewFactory()
    music := factory.MakeRnbMusic()
    music.Play()
    music.Stop()
}
