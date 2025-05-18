package main

import (
    "fmt"
    "github.com/BingguWang/design-pattern-study/creational-pattern/prototype/model"
    "github.com/BingguWang/design-pattern-study/utils"
    "time"
)

/**
原型模式。
    也称为克隆模式
    使你能够复制已有对象， 而又无需使代码依赖它们所属的类。

    平时我们想要复制一个对象的时候， 我们需要知道这个对象是什么类的
    于是复制的过程就和类进行耦合了，有没有办法解耦

    实现：
        实现一个clone接口，也叫作原型接口，这个接口用于克隆对象，当直接创建对象代价大时，可以用此接口创建

    浅拷贝：
        只拷贝对象的基本类型，引用还是原来的。常见的赋值操作其实就是一种浅拷贝
    深拷贝：
        对象中的内容全部拷贝,因为引用也是拷贝的，所以不会影响原来的对象
*/
func main() {
    // 假设我们现在获取一个对象
    object := GetAObject()

    //我们不需要关系该对象的类型,也不需要关心他有什么组成
    // 只需要调用clone方法就可以复制出一份一样的对象，只要他有clone接口我们就可以复制一份
    b := object.Clone()
    fmt.Println(utils.ToJson(object["w"]))
    fmt.Println(utils.ToJson(b["w"]))
    fmt.Println(object["w"] == b["w"])
    fmt.Println(utils.ToJson(object["l"]))
    fmt.Println(utils.ToJson(b["l"]))
    fmt.Println(object["l"] == b["l"])
}

func GetAObject() model.Articles {
    return model.Articles{
        "w": &model.Article{
            Title:     "小王",
            Likes:     10,
            UpdatedAt: time.Now(),
            Desc:      &model.Desc{},
        },
        "l": &model.Article{
            Title:     "小李",
            Likes:     10,
            UpdatedAt: time.Now(),
            Desc:      &model.Desc{},
        },
    }
}
