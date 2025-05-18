package main

import "fmt"

func main() {
	handleccc([]int{1, -2, 3, 4, -9, 7}, 6)
}

func handleccc(nums []int, n int) int {
	presum := make([]int, n+1)
	presum[0] = 0

	for i := 1; i < len(presum); i++ {
		presum[i] = presum[i-1] + nums[i-1]
	}
	fmt.Println(presum)

	// 得到前缀和后，sum(nums[i:j]) = pre[j]-pre[i]
	// i从0开始，那就是要找到一个j的位置,  sum(nums[0:j]) = pre[j]-pre[0]
	// sum(nums[0:j]) 表示0到j-1的和 sum(nums[j:len(nums)])=  pre[len(nums)]-pre[j]
	var res = 0
	for j := 0; j < len(nums); j++ {
		chazhi := abs((presum[len(nums)] - presum[j]) - (presum[j] - presum[0]))
		if chazhi > res {
			res = chazhi
		}
	}
	fmt.Println(res)
	return res
}
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
