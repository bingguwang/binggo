package main

import (
	"fmt"
	"testing"
)

func quickSort2(arr []int) []int {
	if len(arr) <= 1 {
		return arr // 如果数组长度为 0 或 1，直接返回
	}

	// 选择基准值（这里选择第一个元素）
	pivot := arr[0]

	// 定义三个切片：小于、等于、大于基准值的部分
	var less, equal, greater []int

	// 遍历数组，将元素分配到对应的切片中
	for _, value := range arr {
		if value < pivot {
			less = append(less, value)
		} else if value == pivot {
			equal = append(equal, value)
		} else {
			greater = append(greater, value)
		}
	}

	// 递归排序并合并结果
	return append(append(quickSort2(less), equal...), quickSort2(greater)...)
}
func TestNamex(t *testing.T) {
	arr := []int{3, 6, 8, 10, 8, 10, 1, 2, 11, 2, 1}
	fmt.Println("Original array:", arr)

	sortedArr := quickSort2(arr)
	fmt.Println("Sorted array:", sortedArr)

}
