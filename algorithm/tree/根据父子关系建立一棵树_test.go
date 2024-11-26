package tree

import (
	"fmt"
	"testing"
)

type Node struct {
	Id       uint
	ParentId uint
	Name     string
	Value    string
	Children []Node
}

func Tree(node []Node, parent uint) []Node { // 输出有序表
	res := make([]Node, 0)
	// 遍历所有的节点找到 父节点是parent的节点
	for _, v := range node {
		if v.ParentId == parent {
			// 递归向下找，找到父节点是本节点的节点
			v.Children = Tree(node, v.Id)
			res = append(res, v)
		}
	}
	for _, re := range res {
		fmt.Println("id:", re.Id, " name:", re.Name)
	}
	return res
}
func TestCsd(t *testing.T) {

	node := []Node{
		{
			Id:       1,
			ParentId: 0,
			Name:     "root",
			Value:    "1",
		},
		{
			Id:       2,
			ParentId: 1,
			Name:     "root",
			Value:    "1",
		},
		{
			Id:       3,
			ParentId: 2,
			Name:     "root",
			Value:    "1",
		},
		{
			Id:       4,
			ParentId: 2,
			Name:     "root",
			Value:    "1",
		},
		{
			Id:       5,
			ParentId: 1,
			Name:     "root",
			Value:    "1",
		},
	}

	nodes := Tree(node, 0)
	for _, n := range nodes {
		fmt.Println(n.Id)
	}

	fmt.Printf("%+v", nodes)
	PrintTree(nodes)
}

func PrintTree(nodes []Node) {
	fmt.Println(nodes[0])
	for _, node := range nodes {
		fmt.Printf(" Node ID: %v, Parent ID: %v\n", node.Id, node.ParentId)
		if len(node.Children) > 0 {
			PrintTree(node.Children)
		}
	}
}
