package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
*
题目描述
给定一个字符串s，最多只能进行一次变换，返回变换后能得到的最小字符串(按照字典序进行比较)变换规则: 交换字符串中任意两个不同位置的字符。
输入描述
一串小写字母组成的字符串 s, 都是小写字母，长度不超过1000
输出描述
按照要求进行变换得到的最小字符串

这题和 leetcode的最大交换差不多

长度为1000，还是有点长的，可以试一下暴力

	func maximumSwap(num int) int {
	    ans := num
	    s := []byte(strconv.Itoa(num))
	    for i := range s {
	        for j := range s[:i] {
	            s[i], s[j] = s[j], s[i]
	            v, _ := strconv.Atoi(string(s))
	            ans = max(ans, v)
	            s[i], s[j] = s[j], s[i]
	        }
	    }
	    return ans
	}

	func max(a, b int) int {
	    if b > a {
	        return b
	    }
	    return a
	}

比较好的是贪心法
选一个字典序最小的字符，如果有多个那最好选位置靠后的，把这个字符尽可能换到前面的位置
*/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	//ans := baoli(s)
	ans := greedy(s)

	fmt.Println(ans)
}

func baoli(str string) string {
	lst := []rune(str)
	var ans = str
	for i := range lst {
		for j := range lst[:i] {
			lst[i], lst[j] = lst[j], lst[i]
			if ans > string(lst) {
				ans = string(lst)
			}
			lst[i], lst[j] = lst[j], lst[i]
		}
	}
	return ans
}

func greedy(str string) string {

	ans := str
	// 获得字符串s的长度n
	n := len(str)
	// 将s转化为字符切片，方便交换操作
	lst := []rune(str)

	// 构建一个栈，储存原字符串从右往左看遇到的字典序更小的字符的下标
	var stack []int

	// 逆序遍历字符串s
	for i := n - 1; i >= 0; i-- {
		// 如果栈是空栈，或者当前下标i对应的字符lst[i]小于栈顶下标对应的字符lst[stack[len(stack)-1]]
		// 则将坐标i加入stack
		if len(stack) == 0 || lst[i] < lst[stack[len(stack)-1]] { // 比较当前字符与 栈顶下标对应的字符 大小，比栈顶对应的字符还小则入栈
			stack = append(stack, i)
		}
	}
	// 概括一下，上面的循环做的工作其实就是：从后向前遍历字符串，找到一个最小的字符，且最终栈顶所指向的就是这个最小字符的下标

	fmt.Println(stack)
	// 所以栈内的元素是用于交换的潜在元素，而且越靠近栈顶，在进行下面的左到右的遍历时，就越先遇见

	// 正序遍历字符串s
	for i := 0; i < n; i++ {
		// 若出现空栈情况，则退出循环
		if len(stack) == 0 {
			break
		}
		// 如果当前下标i位于栈顶元素stack[len(stack)-1]的左边
		// 则可以进行后续判断
		// 正序遍历我们只考虑那些位于栈顶元素左边的字符，因为我们只会在栈顶元素的左边寻找可以与其交换的字符，
		// 如果 i 已经超过了栈顶元素的下标，说明我们已经遍历到了栈顶元素及其右边的部分，这时我们不再需要考虑与栈顶元素交换，而是应该弹出栈顶元素，继续寻找下一个潜在的交换目标。
		if i < stack[len(stack)-1] {
			// 若当前字符大于栈顶元素对应的字符，则可以进行交换
			if lst[i] > lst[stack[len(stack)-1]] {
				lst[i], lst[stack[len(stack)-1]] = lst[stack[len(stack)-1]], lst[i]
				ans = string(lst)
				break
			}
			// 否则，考虑下一个i，这里的else也可以不写
		} else {
			// 如果当前下标i不位于栈顶元素stack[len(stack)-1]的左边
			// 则弹出栈顶元素，考虑下一个字典序较大但是位于较右位置的字符
			stack = stack[:len(stack)-1]
		}
	}

	// 输出答案
	fmt.Println(ans)
	return ans
}
