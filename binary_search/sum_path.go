/*
*    @author wangchunyan.wang
*    @date 1/20/21 5:51 PM
 */

package binarysearch

/*
给定一个二叉树和一个数字n，判断二叉树中是否有一个路径上的数字之和等于给定的数字n
For example: Given the below binary tree and sum = 22,
      5
     / \
    4   8
   /   / \
  11  13  4
 /  \      \
7    2      1
return true, as there exist a root-to-leaf path 5->4->11->2 which sum is 22.
*/

// Node 二叉树节点
type Node struct {
	Value int64
	Left  *Node
	Right *Node
}

// ExistPathBySum 是否存在和为指定数字的路径
func ExistPathBySum(root *Node, sum int64) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && root.Value == sum {
		return true
	}

	return ExistPathBySum(root.Left, sum-root.Value) || ExistPathBySum(root.Right, sum-root.Value)
}

// ExistPathBySum2 是否存在和为指定数字的路径
func ExistPathBySum2(root *Node, sum, target int64) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && (sum+root.Value) == target {
		return true
	}
	return ExistPathBySum2(root.Left, sum+root.Value, target) || ExistPathBySum2(root.Right, sum+root.Value, target)
}
