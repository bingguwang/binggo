package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// 读取输入字符串s和前缀pre
	scanner.Scan()
	s := scanner.Text()
	scanner.Scan()
	pre := scanner.Text()
	fmt.Println(pre)

	// 初始化列表lst用于存放所有单词
	lst := []string{""}

	// 遍历s中的所有字符ch
	for _, ch := range s {
		if unicode.IsLetter(ch) {
			// 如果是字母，则加入到lst最后一个元素的末尾，即延长当前单词
			lst[len(lst)-1] += string(ch)
		} else {
			// 如果不是字母，说明遇到一个标点符号，结束当前单词的获取，lst的末尾插入一个新的空字符串""
			// 插入空字符串的目的是因为，上面的分支里是在最后一个单词后面拼接字母
			lst = append(lst, "")
		}
	}
	fmt.Println(lst)
	fmt.Println(len(lst))

	// 使用map去重lst中可能出现的重复单词
	wordSet := make(map[string]struct{})
	for _, word := range lst {
		if word != "" {
			wordSet[word] = struct{}{}
		}
	}
	fmt.Println("去重后--", lst)

	// 将去重后的单词放入切片并进行排序
	lstSorted := make([]string, 0, len(wordSet))
	for word := range wordSet {
		lstSorted = append(lstSorted, word)
	}
	sort.Strings(lstSorted)
	fmt.Println(lstSorted)

	// 初始化答案数组
	ans := []string{}

	// 获得pre的长度，用于切片
	preLength := len(pre)
	// 遍历lstSorted中的每一个单词
	for _, word := range lstSorted {
		// 如果word前preLength个字符的切片等于pre
		// 说明word的前缀是pre，将其加入答案数组ans中
		if len(word) >= preLength && word[:preLength] == pre {
			ans = append(ans, word)
		}
	}

	// 如果ans长度大于0，说明至少存在一个单词的前缀是pre，输出由所有单词组成的字符串
	// 如果ans长度等于0，说明不存在任何一个单词的前缀是pre，返回pre
	if len(ans) > 0 {
		fmt.Println(strings.Join(ans, " "))
	} else {
		fmt.Println(pre)
	}
}
