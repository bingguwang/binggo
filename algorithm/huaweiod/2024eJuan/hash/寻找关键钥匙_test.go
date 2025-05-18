package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
	"unicode"
)

/*
*
小强正在参加《密室逃生》游戏，当前关卡要求找到符合给定密码 κ(升序的不重复小写字母组成)的箱子，并给出箱子编号，箱子编号为 1~N 。
每个箱子中都有一个字符串 s，字符串由大写字母，小写字母，数字，标点符号，空格组成，需要在这些字符串中找出所有的字母，忽路大小写且去重后排列出对应的密码串，并返回匹配密码的箱子席号，
注意:满足条件的箱子不超过1个。
输入描述
第一行为表示密码 k 的字符串
第二行为一系列箱子 boxes，为字符串数组样式，以空格分隔箱子 ㎡ 数量满足 1==10888，代表每一个箱子的字符串 s的长度满足 0<=s.1ength<=58，密码为仅包含小写字母的升序字符串日不存在重复字母，密码K长度满足1<=K.length<= 26
输出描述
返回对应箱子编号，如不存在符合要求的密码箱，则返回-1
补充说明
箱子中字符拼出的字符串与密码的匹配忽略大小写，且要求与密码完全匹配，如密码 abc 匹配 aBc，但是密码 abc 不匹配 abcd

需要记住的几个方法
reflect.DeepEqual
unicode.ToLower
unicode.IsLetter
*/
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	pwd := scanner.Text()
	pwddict := make(map[rune]int, 0)
	for _, ch := range pwd {
		lower := unicode.ToLower(ch)
		pwddict[lower]++
	}
	fmt.Println("pwddict---", pwddict)

	scanner.Scan()
	fields := strings.Fields(scanner.Text())

	for i, v := range fields {
		dict := make(map[rune]int, 0)
		for _, ch := range v {
			if !unicode.IsLetter(ch) {
				continue
			}
			lower := unicode.ToLower(ch)
			dict[lower]++
		}
		fmt.Println("dict--", dict)
		if reflect.DeepEqual(dict, pwddict) {
			fmt.Println("找到")
			fmt.Println(i + 1) // 输出下标
			goto stop
		}
	}
	fmt.Println(-1)
stop:
}
