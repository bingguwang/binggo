package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	s2 := strings.Fields(s.Text())
	nums := []int{}
	for _, v := range s2 {
		i, _ := strconv.Atoi(v)
		nums = append(nums, i)
	}
	s.Scan()
	yuangongshu, _ := strconv.Atoi(s.Text())

	sort.Ints(nums)
	var sum int
	for i := range nums {
		sum += nums[i]
	}

	// 我们需要找出最小的k
	// 先看下k的范围,
	// k最大时，就是所有的任务都一个人完成的时候，此时k = sum(nums)
	//  因为 每个部分的和都不能大于k， 所以k 最小不能比 max(nums)小
	left, right := nums[len(nums)-1], sum+1
	for left < right {
		mid := left + (right-left)/2
		if f(nums, mid, yuangongshu) { // 能否保证所有的分组都满足小于mid
			right = mid
		} else {
			left = mid + 1 // 完成不了，需要加大一下k
		}
	}
	fmt.Println(left)
	// 	fmt.Println(right)
}

/**
分成 n分，没分的和，找出这些和的最大值。怎么分这个最大值尽可能小
	试想一下，有个k，使得每一份的大小都不超过k, 那么最后结果就是要找到最小的这个k
*/
// 于是就是子问题，一直到有n个人， 每个人的任务不能超过的k, 能否完成所有的任务
func dfs(nums []int, work []int, index int, k int) bool {
	// work里存的是当前每个人分配到的任务量
	// index 表示当前要分配的任务的下标
	if len(nums) == index {
		// 所有的任务都分配完了
		return true
	}

	for i := 0; i < len(work); i++ {

		if nums[index]+work[i] <= k { // 没有达到限制
			// 任务可以分给当前员工
			work[i] += nums[index]
			// 进入下一次任务的分配
			if dfs(nums, work, index+1, k) {
				return true // 这次的分配可行
			}
			// 下次分配不成功，表示本次任务的分配不可行， 回溯当前的分配
			work[i] -= nums[index]
		}
	}
	return false // 到这里就表示所有的都尝试过了，并且也回溯后尝试过了，分配都不可行
}

func f(nums []int, k int, yuangongshu int) bool {
	work := make([]int, yuangongshu) // 保存所有人当前分配到的任务
	return dfs(nums, work, 0, k)
}
