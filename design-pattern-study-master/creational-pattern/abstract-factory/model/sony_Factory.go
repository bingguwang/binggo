package model

// SonyFactory 具体的工厂类，这里类似唱片公司
type SonyFactory struct {
}

func (*SonyFactory) MakeRnbMusic() IMusic {
    return NewSonyRnb()
}
func (*SonyFactory) MakeRockMusic() IMusic {
    return NewSonyRock()
}
