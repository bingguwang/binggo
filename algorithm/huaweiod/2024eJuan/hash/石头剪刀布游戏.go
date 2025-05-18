package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

/*
*题目描述
石头剪刀布游戏有 3 种出拳形状:石头、剪刀、布。分别用字母 A,B,C 表示。游戏规则:
1.出拳形状之间的胜负规则如下: A>B;B>c;c〉A;""左边一个字母，表示相对优势形状。右边一个字母，表示相对劣势形状。
2.当本场次中有且仅有一种出拳形状优于其它出拳形状，则该形状的玩家是胜利者。否则认为是平局。
3.当发生平局，没有赢家。有多个胜利者时，同为赢家。
例如 1:三个玩家出拳分别是 A，B，c，由于出现三方优势循环(即没有任何一方优于其它出拳者)，判断为平局。
例如 2:三个玩家，出拳分别是 A，B，出拳A的获胜。
例如 3:个玩家，出拳全部是 A，判为平局。

三种都出现，平局
只出现一种，平局
只出现两种，才会有胜局
*/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	mp := make(map[string][]string)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if len(fields) != 2 {
			break
		}
		name, val := fields[0], fields[1]
		mp[val] = append(mp[val], name) // 出拳为val的用户的有哪些
	}
	if err := scanner.Err(); err != nil {
		return
	}

	if len(mp["A"]) > 0 && len(mp["B"]) > 0 && len(mp["C"]) > 0 {
		fmt.Println("NULL")
		return
	}
	var ans []string
	switch {
	case (len(mp["A"]) > 0 && len(mp["B"]) == 0 && len(mp["C"]) == 0) ||
		(len(mp["A"]) == 0 && len(mp["B"]) > 0 && len(mp["C"]) == 0) ||
		(len(mp["A"]) == 0 && len(mp["B"]) == 0 && len(mp["C"]) > 0):
		fmt.Println("NULL")
		return
	case len(mp["A"]) > 0 && len(mp["B"]) > 0 && len(mp["C"]) == 0:
		ans = mp["A"]
	case len(mp["A"]) > 0 && len(mp["B"]) == 0 && len(mp["C"]) > 0:
		ans = mp["C"]
	case len(mp["A"]) == 0 && len(mp["B"]) > 0 && len(mp["C"]) > 0:
		ans = mp["B"]
	}
	sort.Strings(ans)
	fmt.Println(ans)
}
