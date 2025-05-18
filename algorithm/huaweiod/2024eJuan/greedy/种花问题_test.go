package main

import "fmt"

/*
*
假设有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花不能种植在相邻的地块上，它们会争夺水源，两者都会死去。

给你一个整数数组 flowerbed 表示花坛，由若干 0 和 1 组成，其中 0 表示没种植花，1 表示种植了花。另有一个数 n ，能否在不打破种植规则的情况下种入 n 朵花？能则返回 true ，不能则返回 false 。

示例 1：

输入：flowerbed = [1,0,0,0,1], n = 1
输出：true
示例 2：

输入：flowerbed = [1,0,0,0,1], n = 2
输出：false
*/

// i是0且左右都是0才可以在这种花, 类比一下机房布局的那题
func main() {
	canPlaceFlowers([]int{1, 0}, 1)
	canPlaceFlowers([]int{0, 0, 1, 0, 1}, 1)
	canPlaceFlowers([]int{1, 0, 0, 0, 1, 0, 0}, 1)
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	new := make([]int, 0)
	new = append(new, flowerbed[0])
	new = append(new, flowerbed...)
	new = append(new, flowerbed[len(flowerbed)-1])
	fmt.Println("new----", new)
	i := 1
	var ans int
	for i < len(new)-1 {
		if new[i] == 0 && new[i-1] == 0 && new[i+1] == 0 {
			ans++
			new[i] = 1
		}
		i++
	}
	fmt.Println(i)
	fmt.Println(ans)
	return ans >= n
}
