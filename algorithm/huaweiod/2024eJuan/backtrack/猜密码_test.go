package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()

	arr := strings.Split(s, ",")
	scanner.Scan()
	s2 := scanner.Text()

	length, _ = strconv.Atoi(s2)

	sort.Strings(arr) // 先排好序
	dfss(arr, 0)
}

var (
	ress   [][]string
	pathh  []string
	length int
)

func dfss(arr []string, start int) {
	if length <= len(pathh) {
		tmp := make([]string, len(pathh))
		copy(tmp, pathh)
		ress = append(ress, tmp)
		fmt.Println(tmp)
	}

	if start >= len(arr) {
		return
	}

	for i := start; i < len(arr); i++ {
		pathh = append(pathh, arr[i])
		dfss(arr, i+1)
		pathh = pathh[:len(pathh)-1]
	}

}
