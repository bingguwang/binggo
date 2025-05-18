package main

import "github.com/BingguWang/design-pattern-study/structural-pattern/proxy-pattern/model"

/**
  代理模式

    静态代理:
        源类和代理类都实现一样的接口，通过在代理类的实现接口里附加功能

    动态代理


这么一看，好像装饰器模式和静态代理一模一样，还是有点区别的：
    代理模式自行管理其服务对象的生命周期
    装饰器模式的由客户端控制

下面是静态代理
*/
func main() {
    proxy := model.NewProxy(&model.User{})

    proxy.Login()
}
