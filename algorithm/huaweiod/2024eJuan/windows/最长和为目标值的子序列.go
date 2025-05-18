package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	s2 := strings.Split(s.Text(), ",")
	var nums []int
	for _, v := range s2 {
		i, _ := strconv.Atoi(v)
		nums = append(nums, i)
	}
	s.Scan()
	tar, _ := strconv.Atoi(s.Text())
	fmt.Println(handlsse(nums, tar))
}

func handlsse(nums []int, tar int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}

	if sum < tar {
		return -1
	}
	if sum == tar {
		return len(nums)
	}
	left, right := 0, 0
	var res int
	cursum := 0
	for right < len(nums) {
		cursum += nums[right]
		// fmt.Println("cursum ", cursum)

		if cursum == tar {
			if right-left+1 > res {
				res = right - left + 1
			}
		}

		for cursum > tar {
			// 左指针右移，左边的数一个个剔除，直到不大于目标值
			cursum -= nums[left]
			left++
		}

		if cursum == tar { // 再次检查
			if right-left+1 > res {
				res = right - left + 1
			}
		}
		// 除了刚好等于目标值，还可能小于目标值，之后都是右移右指针
		right++
	}
	return res
}
