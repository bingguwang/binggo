package main

import (
	"fmt"
)

func main() {
	s := "quacqkuquacqkacuqkackuack"

	// 构建一个标记，表示是否出现了异常，初始化为 false 表示没有异常
	isError := false

	// 构建哈希表，"quack"分别对应 01234 一共五个索引
	word := map[rune]int{
		'q': 0,
		'u': 1,
		'a': 2,
		'c': 3,
		'k': 4,
	}

	// cnt 列表：记录 "quack" 中每个字符被叫了几次，即每个字符的出现个数。
	cnt := [5]int{0, 0, 0, 0, 0}

	for _, ch := range s {
		// 获得字符 ch 在列表 cnt 中的索引 idx，即表示 "quack" 对应的 01234 一共五个索引
		if idx, ok := word[ch]; ok {
			// ch 对应的计数 +1
			cnt[idx]++
			// ch 不是 "q"，且前一个字符数目少于当前字符数目，出现错误
			if ch != 'q' && idx > 0 && cnt[idx] > cnt[idx-1] { // 这个判断很有必要哦
				isError = true
				break
			}
			// 遇到 "q"，表示出现新的雁叫，可能可以复用之前的大雁，如果：
			// 1. 之前有某大雁叫过 "k"，那么这只大雁可以复用，cnt 整体 -1，表示可以少算一只大雁
			// 2. 之前没有大雁叫过 "k"，那么无法进行大雁的复用，无需做任何操作
			// cnt[4] 表示之前叫过 "k" 的大雁的个数，
			// cnt[4] >= 1 即表示，存在某大雁叫过 "k"，这只大雁可以拿来复用
			if ch == 'q' && cnt[4] >= 1 {
				fmt.Println("发生复用")
				for i := range cnt {
					cnt[i]--
				}
				fmt.Println(cnt)
			}
			//fmt.Println(cnt)
		} else {
			// 如果遇到非法字符，直接标记为错误
			isError = true
			break
		}
	}

	// 排除特殊情况，最终计算结束时，cnt 中的元素应该值相等
	// 如果存在不相等的元素，则说明各个字符的总数不一致，出现异常
	for i := 1; i < len(cnt); i++ {
		if cnt[i] != cnt[0] {
			isError = true
			break
		}
	}

	// 如果 isError 为 true，说明出现了异常，输出 -1
	// 否则最终 cnt 中所有计数一致，这个数即为所需要的大雁的个数，输出之
	if isError {
		fmt.Println(-1)
	} else {
		fmt.Println(cnt[0])
	}
}
