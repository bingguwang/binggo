package main

import "github.com/BingguWang/design-pattern-study/creational-pattern/builder/model"

/**
生成器模式：
    也叫创建者模式，builder模式
    该模式允许你
        使用相同的创建代码生成不同类型和形式的对象


    假设现在有一个复杂的对象，有很多的参数
    构建对象的时候，有些参数不是我们需要的，而我们必须创建的时候也填他们，这种是很鸡肋的

    生成器模式建议将对象构造代码从产品类中抽取出来， 并将其放在一个名为生成器的独立对象中。

*/

func main() {
    builder := model.NewRnbBuilder()
    director := model.NewDirector(builder)

    music := director.Build()
    music.Play()
}
