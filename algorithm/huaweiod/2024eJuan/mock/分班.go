package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func strhandle(str string) bool {
	fmt.Println(str)
	split := strings.Split(str, "/")
	if split[1] == "Y" {
		return true
	}
	return false
}
func nohandle(i int, strs []string) int {
	split := strings.Split(strs[i], "/")
	v, _ := strconv.Atoi(split[0])
	return v
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	s2 := strings.Fields(s.Text())

	//dp[i]=true 表示第i个小朋友在1班
	//dp[i]=false 表示第i个小朋友在2班
	// i是Y dp[i] = dp[i-1]
	// i是N dp[i] = !dp[i-1]

	dp := make([]bool, len(s2))
	dp[0] = true

	for i := 1; i < len(s2); i++ {
		if strhandle(s2[i]) {
			dp[i] = dp[i-1]
		} else {
			dp[i] = !dp[i-1]
		}
	}
	var class1 = make([]int, 0)
	var class2 = make([]int, 0)
	for i := 0; i < len(dp); i++ {
		if dp[i] {
			i2 := nohandle(i, s2)
			class1 = append(class1, i2)
		} else {
			i2 := nohandle(i, s2)
			class2 = append(class2, i2)
		}
	}
	fmt.Println(class1)
	fmt.Println(class2)
	sort.Ints(class1)
	sort.Ints(class2)
}
