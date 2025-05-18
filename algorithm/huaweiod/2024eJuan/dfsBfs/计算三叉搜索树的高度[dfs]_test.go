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
定义构造三又搜索树规则如下:
每个节点都存有一个数，当插入一个新的数时，从 根节点只向下寻找，直到找到一个合适的空节点插入。查找的规则是:
1.如果数小于节点的数减去 500则将数插入节点的左子树2.如果数大于节点的数加上 500则将数插入节点的右子树
3.否则，将数插入节点的中子树
给你一系列数，请按以上规则，按顺序将数插入树中，构建出一棵三叉搜索树，最后输出树的高度,

输入描述
第-行为一个数 N，表示有 N个数，1<=N<= 10000第二行为 N 个空格分隔的整数，每个数的范围为[1，10000]

输出描述
输出树的高度(根节点的高度为1)
*/
type node struct {
	left  *node
	mid   *node
	right *node
	val   int
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	list := make([]int, n)
	scanner.Scan()
	fields := strings.Fields(scanner.Text())
	for i, v := range fields {
		list[i], _ = strconv.Atoi(v)
	}
	root := &node{val: list[0]}
	for _, v := range list[1:] {
		build(v, root)
	}

	dfsCalHeight(root, 1)
	fmt.Println(ans)
}

func build(val int, nod *node) {
	if val < nod.val-500 { // val要放到当前节点的左子树里
		if nod.left == nil {
			nod.left = &node{val: val}
		} else {
			build(val, nod.left) // 递归当前节点变为左子树
		}
	} else if val > nod.val+500 { // val要放到当前节点的右子树里
		if nod.right == nil {
			nod.right = &node{val: val}
		} else {
			build(val, nod.right)
		}
	} else {
		if nod.mid == nil { // val要放到当前节点的中子树里
			nod.mid = &node{val: val}
		} else {
			build(val, nod.mid)
		}
	}
}

// dfs深度遍历递归计算树高度
var ans int

// 计算树高，应该自然想到dfs
func dfsCalHeight(node *node, curDepth int) {
	if node == nil {
		return
	}
	if curDepth > ans {
		ans = curDepth
	}
	dfsCalHeight(node.left, curDepth+1)
	dfsCalHeight(node.mid, curDepth+1)
	dfsCalHeight(node.right, curDepth+1)
}
