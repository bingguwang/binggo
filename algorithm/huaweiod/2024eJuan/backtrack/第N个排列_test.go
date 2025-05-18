package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	s2 := s.Text()

	n, _ := strconv.Atoi(s2)
	res, path = make([][]int, 0), make([]int, 0)
	used = make([]bool, n)
	dfsss(n, used)
}

var (
	used []bool
	res  [][]int
	path []int
)

func dfsss(n int, used []bool) {
	if n == len(path) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		fmt.Println(tmp)
		return
	}

	for i := 1; i <= n; i++ {
		if !used[i-1] {
			used[i-1] = true
			path = append(path, i)
			dfsss(n, used)
			used[i-1] = false
			path = path[:len(path)-1]
		}
	}
}
