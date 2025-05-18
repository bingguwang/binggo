package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	s3 := strings.Fields(s.Text())
	var arr []int
	for _, v := range s3 {
		t, _ := strconv.Atoi(v)
		arr = append(arr, t)
	}
	i := handle(arr)
	fmt.Println(i)
	// 第一步最远范围是
	step := int(math.Ceil(float64(len(arr)) / 2))
	var res = 10000
	for i := 1; i < step; i++ { // 第一步可能的长度都遍历一下
		//  走完第一步之后，索引位置是i
		tmp := handle(arr[i:])
		fmt.Println("第一步走 ", i, "步时: ", tmp)
		if tmp != -1 {
			if res > tmp {
				res = tmp
			}
		}
	}
	if res == 10000 {
		fmt.Println(-1)
		return
	}
	fmt.Println(res)
}

// leetcode的跳跃游戏II的方法
func handle(nums []int) int {
	currange := 0
	nextrange := 0
	var res int
	var flag bool
	for i := 0; i < len(nums); i++ {
		nextrange = max(nextrange, nums[i]+i)
		if i == currange {
			res++
			currange = nextrange
			if currange == len(nums)-1 {
				flag = true
				break
			}
		}
	}
	if !flag {
		return -1
	}
	return res
}
func max(s, b int) int {
	if s > b {
		return s
	}
	return b
}
