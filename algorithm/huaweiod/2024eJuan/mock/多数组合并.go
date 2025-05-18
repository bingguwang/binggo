package main

// 传入的是所有的数组组成的集合，各数组的长度是不一定相同的
func duoShuZuHeBing(arr [][]int, k int) []int {
	var res []int
	flag := true
	idx := 0
	for flag {
		flag = false
		for _, lst := range arr {
			if idx < len(lst) {
				end := idx + k
				if end > len(lst) {
					end = len(lst)
				}
				res = append(res, lst[idx:end]...)
				flag = true
			}
		}
		// 更新起始点为分割点的位置
		idx += k
	}
	return res
}
