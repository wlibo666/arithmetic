package leetcode

import (
	"fmt"

	"../queue"
)

/*
1、先序遍历二叉树顺序(根节点放在最前面)：根节点 –> 左子树 –> 右子树，即先访问根节点，然后是左子树，最后是右子树。
上图中二叉树的前序遍历结果为：0 -> 1 -> 3 -> 4 -> 2 -> 5 -> 6

2、中序遍历二叉树顺序(根节点放在中间)：左子树 –> 根节点 –> 右子树，即先访问左子树，然后是根节点，最后是右子树。
上图中二叉树的中序遍历结果为：3 -> 1 -> 4 -> 0 -> 5 -> 2 -> 6

3、后续遍历二叉树顺序(根节点放在最后面)：左子树 –> 右子树 –> 根节点，即先访问左子树，然后是右子树，最后是根节点。
上图中二叉树的后序遍历结果为：3 -> 4 -> 1 -> 5 -> 6 -> 2 -> 0

4、层序遍历二叉树顺序：从最顶层的节点开始，从左往右依次遍历，之后转到第二层，继续从左往右遍历，持续循环，直到所有节点都遍历完成
上图中二叉树的层序遍历结果为：0 -> 1 -> 2 -> 3 -> 4 -> 5 -> 6

我们如何将二叉树一层一层的遍历输出？其实在这里我们要借助一种数据结构来完成：队列。
我们都知道，队列是一种先进先出的数据结构，我们可以先将整颗二叉树的根节点加入队尾，然后循环出队，
每次读取对头元素输出并且将队头元素出队，然后将这个输出的元素节点的的左右子树分别依次加入队尾，重复这个循环，
知道队列为空的时候结束输出。那么整个二叉树就被我们采用层序遍历的思想输出来了

二叉搜索树 左节点的值小于父节点，父节点值小于右节点

*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreOrderParse(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("node value:%d\n", root.Val)
	PreOrderParse(root.Left)
	PreOrderParse(root.Right)
}

func MedOrderParse(root *TreeNode) {
	if root == nil {
		return
	}
	MedOrderParse(root.Left)
	fmt.Printf("node value:%d\n", root.Val)
	MedOrderParse(root.Right)
}

func PastOrderParse(root *TreeNode) {
	if root == nil {
		return
	}
	MedOrderParse(root.Left)
	MedOrderParse(root.Right)
	fmt.Printf("node value:%d\n", root.Val)
}

// 层序遍历
func SequenceParse(root *TreeNode) {
	queue := queue.Init()
	if root == nil {
		return
	}
	queue.Put(root)

	for !queue.Empty() {
		v := queue.Pop()
		tmpV := v.(*TreeNode)
		fmt.Printf("node value:%d\n", tmpV.Val)

		if tmpV.Left != nil {
			queue.Put(tmpV.Left)
		}
		if tmpV.Right != nil {
			queue.Put(tmpV.Right)
		}
	}
}

// 层序 从右侧开始打印
func SequenceRightParse(root *TreeNode) {
	queue := queue.Init()
	if root == nil {
		return
	}
	queue.Put(root)

	for !queue.Empty() {
		v := queue.Pop()
		tmpV := v.(*TreeNode)
		fmt.Printf("node value:%d\n", tmpV.Val)
		if tmpV.Right != nil {
			queue.Put(tmpV.Right)
		}
		if tmpV.Left != nil {
			queue.Put(tmpV.Left)
		}
	}
}

//
func HasPathSum(node *TreeNode, sum int) bool {
	if node == nil {
		return false
	}
	if node.Left == nil && node.Right == nil && node.Val == sum {
		fmt.Printf("found end node:%d\n", node.Val)
		return true
	}
	var left = sum - node.Val
	return HasPathSum(node.Left, left) || HasPathSum(node.Right, left)
}

func OrderParseTest() {
	node0 := &TreeNode{Val: 0}
	node1 := &TreeNode{Val: 1}
	node2 := &TreeNode{Val: 2}
	node3 := &TreeNode{Val: 3}
	node4 := &TreeNode{Val: 4}
	node5 := &TreeNode{Val: 5}
	node6 := &TreeNode{Val: 6}

	node0.Left = node1
	node0.Right = node2
	node1.Left = node3
	node1.Right = node4
	node2.Left = node5
	node2.Right = node6

	fmt.Printf("PreOrderParse\n")
	PreOrderParse(node0)

	fmt.Printf("MedOrderParse\n")
	MedOrderParse(node0)

	fmt.Printf("PastOrderParse\n")
	PastOrderParse(node0)

	fmt.Printf("SequenceParse\n")
	SequenceParse(node0)

	fmt.Printf("SequenceRightParse\n")
	SequenceRightParse(node0)

	HasPathSum(node0, 5)
}
