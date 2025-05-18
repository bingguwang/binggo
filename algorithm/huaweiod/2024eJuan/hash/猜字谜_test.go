package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

/*
conection
connection,today
*/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputs := strings.Split(scanner.Text(), ",")
	scanner.Scan()
	database := strings.Split(scanner.Text(), ",")
	//fmt.Println(inputs)
	//fmt.Println(database)

	mpdatabase := make(map[string]string)
	for _, v := range database {
		s := modifystring(v)
		//fmt.Println(s)
		mpdatabase[s] = v
	}
	var ans []string
	for _, v := range inputs {
		s := modifystring(v)
		if k, b := mpdatabase[s]; b {
			ans = append(ans, k)
		} else {
			ans = append(ans, "not found")
		}
	}
	fmt.Println(strings.Join(ans, ","))
}

// 字符串去重并且排序
func modifystring(str string) string {
	mp := make(map[rune]bool, 0)
	for _, v := range str {
		mp[v] = true
	}
	var arr = make([]rune, 0)
	for k, _ := range mp {
		arr = append(arr, k)
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	return string(arr)
}
