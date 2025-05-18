package model

// 简单工厂其实不算真正意义上的工厂模式

type IMusicFactory interface {
    MakeRnbMusic() IMusic // 工厂方法,工厂方法返回的是产品
    MakeRockMusic() IMusic
}
type Factory struct {
}

func (f *Factory) MakeRnbMusic() IMusic {
    return NewRnb()
}

func (f *Factory) MakeRockMusic() IMusic {
    return NewRock()
}
func NewFactory() *Factory {
    return &Factory{}
}
