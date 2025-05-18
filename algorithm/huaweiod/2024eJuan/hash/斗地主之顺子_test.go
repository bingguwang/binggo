package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
*
题目描述
在斗地主Q扑克牌游戏中，扑克牌由小到大的顺序为: 3,4,5,6,7,8,9,10,,Q,K,A,2.
玩家可以出的扑克牌阵型有:单张、对子顺子、飞机、炸弹等。其中顺子的出牌规则为:由至少5张由小到大连续递增的扑克牌组成，且不能包含2。
例如:(3,4,5,6,7}、(3,4,5,6,7,8,9,18,,Q,K,A}都是有效的顺子;而 {,Q,K,A,2}、{2,3,4,5,6}、{3,4,5,6}、{3,4,5,6,8}等都不是顺子。给定一个包含 13 张牌的数组，如果有满足出牌规则的顺子，请输出顺子。
如果存在多个顺子，请每行输出一个顺子，且需要按顺子的第一张牌的大小(必须从小到大)依次输出。如果没有满足出牌规则的顺子请输出 N'o。

每一张牌只可以使用一次，但如果能够凑出多个顺子需要尽量去使用,
当出现多个顺子的起始位置相等的时候，应该先输出长度更长的顺子

*/

/*
*
这个和恢复数字序列有点类似，但有区别，因为这里不知道顺子的长度，而那道题里是自己输入区间的长度的
*/

var next_card_dic = map[string]string{
	"3":  "4",
	"4":  "5",
	"5":  "6",
	"6":  "7",
	"7":  "8",
	"8":  "9",
	"9":  "10",
	"10": "J",
	"J":  "Q",
	"Q":  "K",
	"K":  "A",
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	arr := strings.Fields(scanner.Text())

	cardCnt := make(map[string]int)
	for _, v := range arr {
		cardCnt[v]++
	}

	ans := make([][]string, 0)
	start := 3 // 顺子的第一张牌最小是3，最大是10,
	for start <= 10 {
		flag := check(start, cardCnt, &ans)
		if !flag { // 只有不存在顺子的时候才换掉start，因为找到顺子的时候，可能这个start还有顺子
			start++
		}
	}
	fmt.Println("ans--", ans)
	fmt.Println("cardCnt--", cardCnt)

	// 因为check里找出的是长度是5 的顺子，但其实顺子长度可能不只有5，check里是为了尽可能多组顺子
	// 所以还需要进行的顺子的延长，看能不能用剩下的牌来延长顺子
	for i, v := range ans {
		ans[i] = extend(v, cardCnt)
	}
	for _, v := range ans {
		fmt.Println(v)
	}

}
func check(start int, cardCnt map[string]int, ans *[][]string) bool {
	var res []string // 表示顺子
	card := ""
	card += strconv.Itoa(start) //当前牌
	// 找出长度是5 的顺子
	for i := 0; i < 5; i++ {
		if cardCnt[card] > 0 { // 当前牌存在，则拼接到顺子里
			res = append(res, card)
		} else {
			break
		}
		if card != "A" {
			card = next_card_dic[card]
		}
	}
	// 如果res长度是5代表找到了一个顺子
	if len(res) == 5 {
		for _, v := range res {
			cardCnt[v]-- // 使用过的牌更新牌数
		}
		*ans = append(*ans, res)
		// 找到了一个以start为开始的长度是5 的顺子
		return true
	}
	return false
}

func extend(res []string, cardsCnt map[string]int) []string {
	last := res[len(res)-1]
	for last != "A" && cardsCnt[next_card_dic[last]] > 0 { // 牌库里有可用的牌
		// 拼接到res,并且更新剩下的牌库
		last = next_card_dic[last]
		cardsCnt[last]--
		res = append(res, last)
	}
	fmt.Println("拼接后--", res)
	return res
}
