package od

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/**
2
20CNY53fen
53HKD87cents
*/
func main() {
	var n int
	fmt.Scan(&n)

	var str []string
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		scanner.Scan()
		text := scanner.Text()
		str = append(str, text)
	}
	var sum float64
	for i := 0; i < len(str); i++ {
		s := str[i]
		sum += getValAndUnit(s)
	}
	fmt.Println(int(sum))
}

func isDigit2(b byte) bool {
	return b <= '9' && b >= '0'
}

// 处理没每个字符串
func getValAndUnit(str string) float64 {
	var i, j int
	fmt.Println(str)
	var vals []float64
	var units []string
	// 切割字符串。把值和单位都切割出来
	for i < len(str) {
		// i 不是数字，i-1 是数字，str[j:i]是值
		if !isDigit2(str[i]) && i-1 > 0 && isDigit2(str[i-1]) {
			float, _ := strconv.ParseFloat(str[j:i], 64)
			vals = append(vals, float)
			j = i
		}
		// i 是数字，i-1不是，str[j:i]是单位
		if isDigit2(str[i]) && i-1 > 0 && !isDigit2(str[i-1]) {
			fmt.Println(str[j:i])
			units = append(units, str[j:i])
			j = i
		}
		i++
	}
	// 最后一个是单位，别漏了
	units = append(units, str[j:i])
	var res float64
	for i := 0; i < len(units); i++ {
		res += com(vals[i], units[i])
	}
	return res
}
func com(val float64, unit string) float64 {
	switch unit {
	case "CNY":
		return val * 100
	case "JPY":
		return val * 10000 / 1825
	case "HKD":
		return val * 10000 / 123
	case "EUR":
		return val * 10000 / 14
	case "GBP":
		return val * 10000 / 12
	case "fen":
		return val
	case "cents":
		return val * 100 / 123
	case "sen":
		return val * 100 / 1825
	case "eurocents":
		return val * 100 / 14
	case "pence":
		return val * 100 / 12
	}
	return 0
}
