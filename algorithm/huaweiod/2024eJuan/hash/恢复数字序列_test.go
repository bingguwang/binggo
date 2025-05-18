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
对于一个连续正整数组成的序列，可以将其拼接成一个字符串，再将字符串里的部分字符打乱顺序。如序列8910 11 12，拼接成的字符串为 89101112，打乱一部分字符后得到 90811211，原来的正整数 18 就被拆成了。和1。 现给定一个按如上规则得到的打乱字符的字符串，请将其还原成连续正整数序列，并输出序列中最小的数字。
输入描述
输入一行，为打乱字符的字符串和正整数序列的长度，两者间用空格分隔，字符串长度Q不超过 200，正整数不超过 1000，保证输入可以还原成唯一序列。
输出描述
输出一个数字，为序列中最小的数字。

输入
19801211 5
输出
8
*/
func main() {
	// 因为是连续的正整数，且知道正整数范围，可以用暴力的方式的解决

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fields := strings.Fields(scanner.Text())

	n, _ := strconv.Atoi(fields[1])
	str := fields[0]
	mp := make(map[rune]int)
	for _, v := range str {
		mp[v]++
	}
	fmt.Println(mp)

	for i := 1; i < 1001 && n+i-1 < 1001; i++ {
		fmt.Println("i取:", i)
		var strs = ""
		for r := i; r <= i+n-1; r++ {
			s := strconv.Itoa(r)
			strs += s // 把数字转为字符串拼接进去
		}
		// 统计各个数字出现的次数
		tmpmap := make(map[rune]int)
		for _, v := range strs {
			tmpmap[v]++
		}
		//if reflect.DeepEqual(tmpmap, mp) {
		if isEqualMap(tmpmap, mp) {
			fmt.Println("结果是---", i)
			break
		}

	}
}

func isEqualMap(a, b map[rune]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if val, ok := b[k]; !ok || val != v {
			return false
		}
	}
	return true
}
