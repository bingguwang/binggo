package main

import "github.com/BingguWang/design-pattern-study/structural-pattern/bridge-pattern/model"

/**
  桥接模式
        抽象和实现分离，扩展能力强。一般在项目早期使用

        抽象部分 （也被称为接口） 是一些实体的高阶控制层。 该层自身不完成任何具体的工作， 它需要将工作委派给实现部分层 （也被称为平台）。

        这里的dsn和database是抽象
        这里的实现是remote和local， mysql和oracle


      比如我们现在有两个抽象, 形状和颜色，然后我们用聚合的关系来联系形状和颜色。这样我们就可以得到:
        褐色的正方形
        绿色的圆形
        白色的三角形....

    jdbc里的驱动就是用了桥接模式
    可以用于拆分较大较复杂的类
    由于抽象与实现分离，所以可以通过在独立的维度上扩展类
*/
func main() {
    var database model.Database

    // 使用本地连接， 使用oracle驱动
    local := &model.Local{}
    oracle := &model.Oracle{}
    oracle.SetD(local)

    // 配置驱动
    database = oracle

    // 数据库连接
    database.Connect()

}
