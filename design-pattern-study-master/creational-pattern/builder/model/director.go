package model

/**
  可以看到创建者模式有点类似工厂模式，他们最大的区别就是在这个director里

  建造者模式重点关注如何分步生成复杂对象
  抽象工厂专注于生产的对象，会马上返回产品
  而建造者模式允许你在获取产品前执行一些额外构造步骤
*/

type Director struct {
    builder IBuilder
}

func NewDirector(b IBuilder) *Director {
    return &Director{
        builder: b,
    }
}

func (d *Director) Build() *Music {
    music := &Music{}
    d.builder.BuildMusicStep1(music)
    d.builder.BuildMusicStep2(music)
    return music
}
