package od

import (
	"fmt"
	"testing"
)

func TestSadsa(t *testing.T) {
	//str := "AASW"
	//str := "AASAAAAA"
	str := "AASSS"
	l := len(str)
	avg := l / 4

	mp := make(map[byte]int)
	for i := 0; i < len(str); i++ {
		mp[str[i]]++
	}
	fmt.Println(mp)
	// targetChars 储存超出平均次数的字符， targetNum为总共的超出的数量
	targetChars := make(map[byte]int)
	targetNum := 0
	for k, v := range mp {
		if v > avg { // 超出平均次数的字符,存起来
			targetChars[k]++
			targetNum += v - avg
		}
	}
	if targetNum == 0 { //如果没有字母超出平均 则本身为完美走位
		fmt.Println(0)
		return
	}
	fmt.Println(targetNum)
	fmt.Println(targetChars)
	// 双指针
	left, right := 0, 0
	minLength := len(str)
	for right < len(str) {
		_, ok := targetChars[str[right]]
		if ok {
			targetNum--
		}
		if targetNum == 0 {
			fmt.Println("可行区间：", left, " ", right)
			for targetNum == 0 {
				_, ok := targetChars[str[left]]
				if ok {
					targetNum++
				}
				left++
			}
			currLength := len(str[left-1 : right+1])
			fmt.Println(str[left-1 : right+1])
			fmt.Println(left-1, " ", right)
			minLength = gm(currLength, minLength)
		}
		right++
	}
	fmt.Println(minLength)
}
func gm(i, j int) int {
	if i < j {
		return i
	}
	return j
}
