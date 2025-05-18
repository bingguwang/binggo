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

怎么贪心：
只要缓存的后面都不用金币了
某个文件是否缓存？
缓存、扫描一次+缓存一次 =  m+m
不缓存、  m * 文件数
比较这两个大小就知道是否缓存此文件了
*/
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text()) // 缓存一次需要的金币
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

	// 准备工作结束
	costscanmp := make(map[int]int)  // 文件全部扫描花费的金币数，key是文件的id
	costcachemp := make(map[int]int) // 文件缓存需要花费的金币数 扫描一次+缓存一次
	for i, id := range ids {
		if _, ok := costscanmp[id]; ok { // 已存在，则累加花费就行
			costscanmp[id] += costs[i]
		} else {
			costscanmp[id] = costs[i]
			costcachemp[id] += costs[i] + n
		}
	}

	fmt.Println(costscanmp)  // 不使用缓存的时候，每个文件需要的金币数
	fmt.Println(costcachemp) //  使用缓存的时候，每个文件需要的金币数

	var ans int
	for k, v := range costscanmp {
		ans += min(v, costcachemp[k]) // 每个文件选择或不选择使用缓存，选择消耗小的那种
	}
	fmt.Println("ans:", ans)

}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
