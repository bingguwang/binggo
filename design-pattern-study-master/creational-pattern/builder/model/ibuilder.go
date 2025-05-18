package model

type IBuilder interface {
    BuildMusicStep1(music *Music)
    BuildMusicStep2(music *Music)
}

func GetBuilder(tp string) IBuilder {
    switch tp {
    case "rnb":
        return NewRnbBuilder()
    case "rock":
        return NewRockBuilder()
    }
    return nil
}
