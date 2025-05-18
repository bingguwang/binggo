package main

import (
	"fmt"
	"math"
)

/*
*跳房子，也叫跳飞机，是一种世界性的儿童游戏。 游戏参与者需要分多个回合按顺序跳到第1格直到房子的最后一格
跳房子的过程中，可以向前跳，也可以向后跳。假设房子的总格数是count ，小红每回合可能连续跳的步教都放在数组steps 中，请问数组中是否有一种步数的组合，可以让小红两个回合跳到最后一格?如果有，请输出索引和最小的步数组合。
注意:
·数组中的步数可以重复，但数组中的元素不能重复使用，
。提供的数据保证存在满足题目要求的组合，且索引和最小的步数组合是唯一的.
输入描述
第一行输入为每回合可能连续跳的步数，它是整数数组类型，第二行输入为房子总格数 count ，它是 int 整数类型Q
输出描述
返回索引和最小的满足要求的步数组合(顺序保持 steps 中原有顺序)

和leetcode的两数之和一模一样
*/

func main() {
	handle(7, []int{1, 4, 5, 2})
}

func handle(tar int, nums []int) [2]int {
	mp := make(map[int]int, 0) // 存着每个数的下标

	var idxsum = math.MaxInt
	var res [2]int
	for i := 0; i < len(nums); i++ {
		val := tar - nums[i]
		if idx, ok := mp[val]; ok {
			fmt.Println("b:", idx, "值：", nums[idx])
			fmt.Println("a:", i, "值：", nums[i])
			// return 不能马上返回
			// 先记录下一次结果
			if idxsum > i+idx {
				idxsum = i + idx
				res = [2]int{nums[idx], nums[i]} // v是先出现的
			}
		} else {
			if _, exist := mp[nums[i]]; !exist { // 有相同值的时候，保证只存第一个的下标
				mp[nums[i]] = i
			}
		}
	}
	fmt.Println(res)
	return res
}
