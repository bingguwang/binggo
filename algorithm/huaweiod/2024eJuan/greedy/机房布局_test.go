package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
*
小明正在规划一
个大型数据中心机房，为了使得机柜上的机器都能正常满负荷工作，需要确保在每个机柜边上至少要有一个电箱。
为了简化题目，假设这个机房是一整排，"表示机柜，ī表示间隔，请你返回这整排机柜，至少需要多少个电箱。 如果无解请返回 -1
*/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	cabinets := "A" + input[:len(input)-1] + "A" // 在cabinets字符串前后添加字符"A"

	n := len(cabinets) // 获得cabinets字符串的长度
	index := 1         // 遍历索引从1开始
	ans := 0           // 设置答案为0， 电箱的数目

	// 进行while循环，退出条件为index已经到达倒数第二个位置
	for index < n-1 {
		// index是个空位，直接跳过
		if cabinets[index] == 'I' {
			index++
		} else if cabinets[index] == 'M' { // index是个机器
			// index后面有空位，则在 index+1 装电箱，index前进3步
			if cabinets[index+1] == 'I' {
				ans++
				index += 3
			} else if cabinets[index+1] != 'I' && cabinets[index-1] == 'I' {
				// index后面没空位，前面有空位。电箱装在 i-1处
				ans++
				index++
			} else if cabinets[index+1] != 'I' && cabinets[index-1] != 'I' {
				// index后面没空位，前面也没有空位
				ans = -1
				break
			}
		}
	}

	fmt.Println(ans) // 输出结果
}
