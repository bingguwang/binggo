package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func result(s string) string {
	// 不区分大小写
	s = strings.ToLower(s)
	fmt.Println("全转为小写字母:", s)
	// 统计出来每个字母出现的次数
	count := make(map[rune]int)
	letters := []rune{} // 去除字符串中的非字母
	for _, c := range s {
		if c >= 'a' && c <= 'z' {
			// 统计字符串
			count[c]++
			letters = append(letters, c)
		}
	}
	fmt.Println("每个字符出现的次数:", count)
	fmt.Println(string(letters))
	s = string(letters) + " " // 加空格是为了避免后续的收尾操作
	count[' '] = 1
	// 记录连续字母和非连续字母
	ans := [][]interface{}{}
	// 上一个字母的位置
	pre := rune(s[0])
	// 该字母的连续次数， 初始化为1
	repeat := 1
	// 后续该字母还有count[pre] -= 1
	count[pre]--
	for i := 1; i < len(s); i++ {
		// 当前位置的字母
		cur := rune(s[i])
		// 后续该字母还有count[cur] -= 1个
		count[cur]--
		if cur == pre {
			// 如果当前位置和上一个位置字母相同，就产生连续
			repeat++
		} else {
			// 当前字母与上一个字母不连续，就打断联系
			var val interface{}
			fmt.Println("repeat:", repeat)
			if repeat > 1 {
				val = repeat
			} else {
				val = count[pre] // 上一个字符的计数结束了, 且是一个字母不是连续的一段，count[pre]里存的是后续还有多少个该字符
			}
			fmt.Println("pre:", string(pre), " val:", val)
			ans = append(ans, []interface{}{pre, val})
			pre = cur // 更新
			// 更新pre连续次数为1
			repeat = 1
		}
	}
	fmt.Println("ans:", ans)
	// 母和紧随的数字作为一组进行排序，数字大的在前，数字相同的，则按字母进行排序，字母小的在前。
	sort.Slice(ans, func(i, j int) bool {
		if ans[i][1].(int) != ans[j][1].(int) {
			return ans[i][1].(int) > ans[j][1].(int) // 先排数字
		}
		return ans[i][0].(rune) < ans[j][0].(rune) // 再排字母
	})
	result := ""
	for _, pair := range ans {
		result += fmt.Sprintf("%c%d", pair[0], pair[1])
	}
	return result
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	s2 := s.Text()
	fmt.Println(result(s2))
}
