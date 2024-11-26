package dp

import (
	"fmt"
	"testing"
)

func TestNamase(t *testing.T) {
	BACKSTR("12HHHHA")
}

func BACKSTR(s string) {
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	res := ""
	for i := len(s) - 1; i >= 0; i-- { // 要倒着迭代，假设我们从0开始迭代，当我们在考虑长度为3的子串时，需要使用dp[i+1][j-1]的值，但是dp[i+1][j-1]的值在此之前并没有被计算出来，因此无法使用
		for j := i; j < len(s); j++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else if j-i < 3 { // // 字符相等小于等于3个字符，就是最多只有2个字符
				dp[i][j] = true
			} else {
				dp[i][j] = dp[i+1][j-1] // 因为我们需要先有dp[i+1]的值才能算dp[i],所以i需要倒着迭代
			}
			// 更新最大值
			if dp[i][j] && j-i+1 > len(res) {
				res = s[i : j+1]
			}
		}
	}
	fmt.Println(len(res))
}
