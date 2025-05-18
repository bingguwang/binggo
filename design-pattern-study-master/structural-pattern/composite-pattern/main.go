package main

import "github.com/BingguWang/design-pattern-study/structural-pattern/composite-pattern/model"

/**
  组合模式类似装饰器模式
      区分他们的方式就是：
          组合模式是通过组合的思想来实现的
          装饰器模式是通过聚合的思想来实现的

  聚合比组合更为松耦合，组合的关系比聚合更紧密一些，
  组合有一种部分和整体的概念，子结构一般是树形的，而聚合没有部分与整体的关系，有层级的关系，子结构一般是平级的

  组合模式也是很常见的
*/
func main() {
    o := &model.Order{
        Name: "book",
        Components: []model.Component{
            &model.OrderItem{Name: "aaa"},
            &model.OrderItem{Name: "bbb"},
            &model.OrderItem{Name: "vvv"},
        },
    }
    o.Desc()
}
