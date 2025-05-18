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
	// 读取输入
	scanner := bufio.NewScanner(os.Stdin)

	// 数组长度
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	// 数组
	scanner.Scan()
	line := scanner.Text()
	numsStr := strings.Split(line, " ")
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i], _ = strconv.Atoi(numsStr[i])
	}

	// 最小能力值要求
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	// 对nums进行排序，方便使用头尾相向双指针
	sort.Ints(nums)

	// 如果nums的最小值已经大于m，则每一个人都单独组队，直接输出n
	if nums[0] >= m {
		fmt.Println(n)
		return
	}

	// 否则，实现贪心+双指针算法
	// 设置头尾双指针
	left, right := 0, n-1
	// 初始化答案变量
	ans := 0

	// 循环遍历，退出循环的条件为left和right指向同一个元素
	for left < right {
		// 如果nums[right]本身已经大于等于m，则nums[right]单独组队
		// right递减，ans更新
		if nums[right] >= m {
			right--
			ans++
		} else if nums[right]+nums[left] >= m {
			// 如果当前剩余数字中的最大值和最小值相加大于等于m
			// 则说明nums[right]带得动nums[left]
			// 贪心地令它们进行组队，两个指针向中间移动，ans更新
			left++
			right--
			ans++
		} else {
			// 如果当前剩余数字中的最大值和最小值相加小于m
			// 则说明即使是能力值最大的nums[right]也带不动nums[left]
			// nums[left]不应该去组队，left递增以找到一个更大的能力值
			left++
		}
	}

	// 退出循环时，存在left >= right
	// 可能出现left = right的情况
	// 但由于所有可以单独组队的right都已经被排除
	// 且最终right不可能降为0
	// （因为nums[0] >= m的情况已经在前面的if判断过了）
	// 故此时nums[left] = nums[right]不能单独组队
	// 直接输出ans即为答案
	fmt.Println(ans)
}
