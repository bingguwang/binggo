package tree

/*
*

哈夫曼树

	构建哈夫曼树的步骤大致：

1  建立一个叶子节点的数组
2  从数组里找出值最小的两个节点，新建的节点再继续假如到原数组里
3  我们需要保存每个节点的高度
4  当数组里只剩下一个节点的时候，这个就是根节点
*/

// 节点的结构
type treeNode struct {
	// 左右节点
	left  *treeNode
	right *treeNode
	// 节点的高度
	height int
	// 节点的值
	val int
}

func main() {

}
