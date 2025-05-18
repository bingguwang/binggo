package main

import "github.com/BingguWang/design-pattern-study/creational-pattern/abstract-factory/model"

/**
  抽象工厂
      可以视为“工厂的工厂”

        不再是已产品为单元，而是以产品簇为单位
*/
func main() {
    // 获取工厂
    sonyFactory := model.GetMusicFactory("sony")

    rockMusic := sonyFactory.MakeRockMusic()
    rnbMusic := sonyFactory.MakeRnbMusic()

    rockMusic.Play()
    rnbMusic.Play()

}
