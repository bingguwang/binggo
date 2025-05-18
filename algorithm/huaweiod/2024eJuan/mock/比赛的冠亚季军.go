package main

import (
	"fmt"
)

func main() {
	guanyajun([]int{2, 3, 4, 5})
}

type player struct {
	id    int
	score int
}

func guanyajun(arr []int) {
	// 淘汰的人都保存下来
	var players []player
	for i, v := range arr {
		players = append(players, player{i, v})
	}
	var miss []player // 记录所有被淘汰的人
	var flag bool
	for len(players) > 1 {
		var nextRound []player
		for i := 0; i < len(players); i += 2 {
			if i+1 >= len(players) { //轮空选手，直接晋级
				nextRound = append(nextRound, player{i, players[i].score})
			} else {
				if aWinb(players[i], players[i+1]) {
					// i晋级
					nextRound = append(nextRound, players[i])
					miss = append(miss, players[i+1])
				} else {
					// i+1晋级
					nextRound = append(nextRound, players[i+1])
					miss = append(miss, players[i])
				}
			}
			if len(nextRound)%4 == 0 {
				flag = true // 如果出现只有4个人的战局。注意越往后的比赛人数越少
			}
		}
		// 更新当前比赛队员
		players = nextRound
	}
	// 被冠军和亚军淘汰的人集合
	fmt.Println("被冠军和亚军淘汰的人集合")
	fmt.Println(miss)
	fmt.Println(players[0]) // 最后集合只剩一人就是冠军
	fmt.Println("倒数第一个淘汰者就是亚军--")
	fmt.Println(miss[len(miss)-1])
	if flag {
		// 如果 flag 为 true（剩下四个选手），季军为倒数第二个和倒数第三个失败者中更强的一个。
		fmt.Println("季军是:")
		if aWinb(miss[len(miss)-2], miss[len(miss)-3]) {
			fmt.Println(miss[len(miss)-2])
		} else {
			fmt.Println(miss[len(miss)-3])
		}
	} else {
		// 如果 flag 为 false（剩下三个选手），季军为倒数第二个失败者。
		fmt.Println(miss[len(miss)-2])
	}
}
func aWinb(a, b player) bool {
	return a.score > a.score || a.id == b.id
}
