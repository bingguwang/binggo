package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	for i := 0; i < n; i++ {
		scanner.Scan()
		fields := strings.Fields(scanner.Text())
		fmt.Println(check(fields))
	}

}
func check(days []string) bool {
	var absent = 0
	for _, v := range days {
		if v == "absent" {
			absent++
		}
	}
	// 缺席超过一次
	if absent > 1 {
		return false
	}
	// 连续的迟到早退
	for i := 0; i < len(days)-1; i++ {
		if (days[i] == "late" || days[i] == "leaveearly") && (days[i+1] == "late" || days[i+1] == "leaveearly") {
			return false
		}
	}

	// 任意7天里出现了迟到早退或缺席次数超过3次
	if len(days) >= 7 {
		for i := 0; i < len(days); i++ {
			var count int
			for j := i; j < i+7 && j < len(days); j++ {
				if days[j] == "absent" || days[j] == "leaveearly" || days[j] == "late" {
					count++
				}
			}
			if count > 3 {
				return false
			}
		}
	}

	return true
}

/**
 // 固定滑窗过程：考虑任意连续的 7 天是否存在 3 天以上的迟到/早退/缺席
        cntWin := make(map[string]int)
		// 计算第一个窗口
        for i := 0; i < 7 && i < n; i++ {
                cntWin[lst[i]]++
        }
		fmt.Println(cntWin)
        if cntWin["leaveearly"]+cntWin["late"]+cntWin["absent"] > 3 {
                return "false"
        }

        // 滑窗过程
        for right := 7; right < n; right++ {
                infoRight := lst[right]
                cntWin[infoRight]++
                left := right - 6
                infoLeft := lst[left]
                cntWin[infoLeft]--
                if cntWin[infoLeft] == 0 {
                        delete(cntWin, infoLeft)
                }
                if cntWin["leaveearly"]+cntWin["late"]+cntWin["absent"] > 3 {
                        return "false"
                }
        }

*/
