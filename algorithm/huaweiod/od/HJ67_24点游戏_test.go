package od

import (
	"fmt"
	"math"
	"testing"
)

/**
给出4个1-10的数字，通过加减乘除运算，得到数字为24就算胜利,除法指实数除法运算,运算符仅允许出现在两个数字之间,本题对数字选取顺序无要求，但每个数字仅允许使用一次，且需考虑括号运算
此题允许数字重复，如3 3 4 4为合法输入，此输入一共有两个3，但是每个数字只允许使用一次，则运算过程中两个3都被选取并进行对应的计算操作。

#############################################
思路就是：
游戏的第一步是挑出两个数，算出一个新数替代这两个数。
然后，在三个数中玩 24 点，再挑出两个数，算出一个数替代它们。
然后，在两个数中玩 24 点……
最后只有一个数了， 我们就只要判断这个数是不是等于24就行了


因为是可以使用括号的，所以上面的思路是可行的
#############################################

*/

// 考察 dfs 递归

func dfs(nums []float64) bool {
	if len(nums) == 1 { // 数组只有1个数了
		return math.Abs(nums[0]-24) < 1e-9 // 是24
	}
	flag := false
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ { // 2层for循环枚举出了所有:挑出两个数  的可能
			n1, n2 := nums[i], nums[j]
			newNums := make([]float64, 0, len(nums)) // numNums保存的是两数计算结果替换两数后，新的数组
			for k := 0; k < len(nums); k++ {
				if k != i && k != j {
					newNums = append(newNums, nums[k]) // 先把没被选中的数加到numNums里
				}
			}

			flag = flag || dfs(append(newNums, n1+n2))
			flag = flag || dfs(append(newNums, n1-n2))
			flag = flag || dfs(append(newNums, n2-n1))
			flag = flag || dfs(append(newNums, n1*n2))
			if n1 != 0 {
				flag = flag || dfs(append(newNums, n2/n1))
			}
			if n2 != 0 {
				flag = flag || dfs(append(newNums, n1/n2))
			}

			if flag {
				return true
			}
		}
	}
	return false
}

func TestSD(t *testing.T) {
	s := make([]int, 4)
	for i := 0; i < 4; i++ {
		fmt.Scan(&s[i])
	}
	f := make([]float64, 4)
	for i := 0; i < 4; i++ {
		f[i] = float64(s[i])
	}
	b := dfs(f)
	fmt.Println(b)
}

// --------------------------------------------------- 解法2

func TestJas(t *testing.T) {
	var a [4]int
	for i := 0; i < len(a); i++ {
		fmt.Scan(&a[i])
	}

	var used [4]bool
	var canDo func(int) bool
	canDo = func(ret int) bool {
		if used[0] == true && used[1] == true && used[2] == true && used[3] == true {
			// 四个数都选过了，直接判断计算结果ret是不是24
			return ret == 24
		}
		for i := 0; i < len(a); i++ {
			if !used[i] { // 去剩下的没有访问的数值取
				used[i] = true
				if canDo(ret+a[i]) ||
					canDo(ret-a[i]) ||
					canDo(ret*a[i]) ||
					a[i] != 0 && ret%a[i] == 0 && canDo(ret/a[i]) {
					return true
				}
				// 回溯，恢复场景
				used[i] = false
			}
		}
		return false
	}

	for i := 0; i < len(a); i++ {
		used[i] = true
		if canDo(a[i]) { // 可以找到解
			fmt.Println("true")
			return
		}
		// 否则回溯
		used[i] = false
	}
	fmt.Println("false")
}
