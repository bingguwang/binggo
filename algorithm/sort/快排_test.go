package main

import (
	"fmt"
	"testing"
)

/*
*
时间复杂度
*/
func Test(t *testing.T) {
	a := []int{2, 3, 1, 3, 4, 5}
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)
}
func quickSort(nums []int, left, right int) {
	if left > right {
		return
	}
	i, j := left, right
	p := nums[left] // 选择最左边的为基准值
	for i != j {
		for nums[j] >= p && i < j { // j必须先走，因为有i<j的限制，如果说j后走哦，可能最后范围由于i的右移导致最后有的部分j本来要遍历的但是没遍历到
			j--
		}
		for nums[i] <= p && i < j {
			i++
		}
		if i < j { // 交换i,j处的值
			fmt.Println("i:", i, "j:", j)
			tmp := nums[i]
			nums[i] = nums[j]
			nums[j] = tmp
		}
	}
	// 此时i,j相遇，交换基准值和i的值
	nums[left] = nums[i]
	nums[i] = p

	quickSort(nums, left, i-1)
	quickSort(nums, i+1, right)
}
func TestJksdsqa(t *testing.T) {
	a := []int{5, 2, 22, 3, 4, 8, 21}
	Quik(a, 0, len(a)-1)
	fmt.Println(a)
}
func Quik(a []int, left, right int) {
	if left > right {
		return
	}
	i, j := left, right
	piv := a[left]
	for i != j {
		for a[j] >= piv && i < j {
			fmt.Println("j", j)
			j--
		}
		for a[i] <= piv && i < j {
			fmt.Println("i", i)
			i++
		}
		// 交换ij
		//if i<j {
		tp := a[i]
		a[i] = a[j]
		a[j] = tp
		//}

	}
	// 交换基准和i
	a[left] = a[i]
	a[i] = piv

	Quik(a, left, i-1)
	Quik(a, i+1, right)
}
