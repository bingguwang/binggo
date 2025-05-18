package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	n, _ := strconv.Atoi(s.Text())
	var res float64
	for i := 0; i < n; i++ {
		s.Scan()
		str := s.Text()
		var tmpunit []rune
		num := 0
		for idx, ch := range str {
			if unicode.IsDigit(ch) { // 数字
				num = 10*num + int(ch-'0')
			} else { // 字母
				fmt.Println(string(ch))
				tmpunit = append(tmpunit, ch)

				if idx == len(str)-1 ||
					idx+1 < len(str) && (unicode.IsDigit(rune(str[idx+1]))) { // 单位取完了
					fmt.Println("数:", num)
					fmt.Println("单位:", string(tmpunit))
					res += (float64(num) * exchange(string(tmpunit)))
					num = 0
					tmpunit = tmpunit[0:0]
				}
			}

		}
	}
	fmt.Println(int(res))
}

// 转为fen
func exchange(unit string) float64 {
	switch unit {
	case "CNY":
		return 100.0
	case "JPY":
		return 100.0 / 1825 * 100.0
	case "HKD":
		return 100.0 / 123 * 100.0
	case "EUR":
		return 100.0 / 14 * 100.0
	case "GBP":
		return 100.0 / 12 * 100.0
	case "fen":
		return 1.0
	case "sen":
		return 100.0 / 1825
	case "cents":
		return 100.0 / 123
	case "eurocents":
		return 100.0 / 14
	case "pence":
		return 100.0 / 12
	}
	return 0.0
}
