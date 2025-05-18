package model

type IMusicFactory interface {
    MakeRnbMusic() IMusic
    MakeRockMusic() IMusic
}

func GetMusicFactory(brand string) IMusicFactory {
    switch brand {
    case "sony":
        return new(SonyFactory)
    case "shack":
        return new(ShackFactory)
    }
    return nil
}
