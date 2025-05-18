package main

import "fmt"

func minEatingSpeed(piles []int, h int) int {
	// 最大速度
	var maxspeed int
	for _, v := range piles {
		if maxspeed < v {
			maxspeed = v
		}
	}
	fmt.Println(maxspeed)
	// 最小速度是0

	// 找到最小的速度，使得可以在h小时内吃完
	left, right := 0, maxspeed
	for left+1 < right {
		mid := (left + right) / 2
		var needtime = len(piles) // 吃完需要的时间
		for _, v := range piles {
			needtime += (v - 1) / mid
		}
		if needtime <= h { // 吃得完，而且时间有余，速度可以放慢
			right = mid
		} else {
			left = mid
		}
		// fmt.Println( left, " ", right) // 掌握不好可以答应出来看看
	}
	fmt.Println(left)
	fmt.Println(right)
	return right
}
