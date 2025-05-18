package model

type ShackFactory struct {
}

func (*ShackFactory) MakeRnbMusic() IMusic {
    return nil
}
func (*ShackFactory) MakeRockMusic() IMusic {
    return nil
}
