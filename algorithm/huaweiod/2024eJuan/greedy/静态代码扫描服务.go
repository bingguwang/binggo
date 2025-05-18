package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//restorecost := 5
	// fileno := []int{1, 2, 2, 1, 2, 3, 4}
	// filesize := []int{1, 1, 1, 1, 1, 1, 1}
	//fileno := []int{2, 2, 2, 2, 2, 5, 2, 2, 2}
	//filesize := []int{3, 3, 3, 3, 3, 1, 3, 3, 3}
	//handlestatic(fileno, filesize, restorecost)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	restorecost, _ := strconv.Atoi(scanner.Text()) // 缓存一次需要的金币
	scanner.Scan()
	fields := strings.Fields(scanner.Text())
	ids := make([]int, 0) // id集合
	for _, v := range fields {
		atoi, _ := strconv.Atoi(v)
		ids = append(ids, atoi)
	}
	fmt.Println(ids)
	scanner.Scan()
	fields = strings.Fields(scanner.Text())
	costs := make([]int, 0) // 扫描成本
	for _, v := range fields {
		atoi, _ := strconv.Atoi(v)
		costs = append(costs, atoi)
	}
	fmt.Println(costs)

	handlestatic(ids, costs, restorecost)

}

func handlestatic(fileno []int, filesize []int, restorecost int) {
	// 需要最少的金币
	// 贪心一点，对于次数多的，也就是缓存一次+文件大小 < 出现次数*文件大小的 ,就缓存

	filecnt := make(map[int]int, 0) //  k 文件序号 v文件出现次数
	filemp := make(map[int]int, 0)  // k 文件序号 v文件大小
	for i := 0; i < len(fileno); i++ {
		filecnt[fileno[i]]++
		filemp[fileno[i]] = filesize[i]
	}
	fmt.Println(filecnt)
	fmt.Println(filemp)
	var res int
	for k, v := range filemp {
		if v+restorecost < filecnt[k]*v { // 缓存
			res += v + restorecost
		} else { // 不缓存
			res += filecnt[k] * v
		}
	}
	fmt.Println(res)
}
