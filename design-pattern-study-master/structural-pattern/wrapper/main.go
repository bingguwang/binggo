package main

import (
    "fmt"
    "github.com/BingguWang/design-pattern-study/structural-pattern/wrapper/model"
)

/**
  装饰器模式：
      wrapper
        将对象放入  包含行为的特殊封装对象中   来为  原对象  绑定新的行为。而不用为此去新建一个子类

    可以防止子类数量爆炸

    当你希望在运行时为对象新增额外的行为， 可以使用装饰模式。而不需要去修改对象

    实现：
        基础行为放到基类，新的行为放到装饰器里
        目标对象和装饰器遵循同一接口，

其实核心就是聚合！！！的思想，用聚合代替了继承，更加灵活，继承有严重的耦合性，聚合没有

    很多代码里都可以看到，这个是很常见的模式

*/
func main() {
    mobile := &model.Mobile{}

    mobilewithgame := &model.MobileWithGame{Mobile: mobile}
    mobilewithwechat := &model.MobileWithWeChat{Mobile: mobilewithgame}

    fmt.Println("功能有：")
    mobilewithwechat.Function()
}
