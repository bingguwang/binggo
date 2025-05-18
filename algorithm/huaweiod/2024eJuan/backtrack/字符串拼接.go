package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var ans int

// 回溯的函数
func dfs(cnt map[rune]int, path string, n int) {
	// 如果path长度已经为n，找到了一个答案
	if len(path) == n {
		fmt.Println(path)
		ans++
		return
	}
	// 横向遍历，考虑字符k以及其剩余个数v
	for k, v := range cnt {
		// 如果剩余个数为0，或者k和path前一个字符一样，则直接跳过k
		if v == 0 || (len(path) > 0 && k == rune(path[len(path)-1])) {
			continue
		}
		// 状态更新：选择k，所以k的剩余个数-1
		cnt[k]--
		// 回溯：path的末尾要加上k这个字符来作为新的path
		dfs(cnt, path+string(k), n)
		// 回滚：把选择的这个k还回去，所以k的剩余个数+1
		cnt[k]++
	}
}

func main() {
	// 使用bufio读取输入
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	parts := strings.Split(input, " ")
	s := parts[0]
	n, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println(0)
		return
	}

	// 统计每个字符出现的次数
	cnt := make(map[rune]int)
	for _, ch := range s {
		cnt[ch]++
	}

	// 检查是否所有字符都是小写字母
	valid := true
	for ch := range cnt {
		if !unicode.IsLower(ch) {
			valid = false
			break
		}
	}

	if valid {
		ans = 0
		dfs(cnt, "", n)
		fmt.Println(ans)
	} else {
		fmt.Println(0)
	}
}
