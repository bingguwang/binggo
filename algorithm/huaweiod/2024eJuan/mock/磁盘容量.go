package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"unicode"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	n, _ := strconv.Atoi(s.Text())
	var strs []string
	for i := 0; i < n; i++ {
		s.Scan()
		s1 := s.Text()
		str2cap(s1)
		strs = append(strs, s1)
	}

	sort.Slice(strs, func(i, j int) bool {
		a, b := strs[i], strs[j]
		if str2cap(a) == str2cap(b) {
			return i < j
		}
		return str2cap(a) < str2cap(b)
	})
	for i := 0; i < n; i++ {
		fmt.Println(strs[i])
	}
}
func str2cap(str string) int { // 单位mb
	var res int
	i := 0
	var shuzi string
	for i < len(str) {
		if unicode.IsDigit(rune(str[i])) {
			shuzi += string(str[i])
			i++
		} else {
			if len(shuzi) != 0 { // 前面有数了
				val, _ := strconv.Atoi(shuzi)
				if str[i] == 'M' {
					res += val
				} else if str[i] == 'G' {
					res += val * 1024
				} else if str[i] == 'T' {
					res += val * 1024 * 1024
				}
				shuzi = shuzi[0:0]
				i++
			}
		}
	}
	return res
}
