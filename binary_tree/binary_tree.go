package binary_tree

type BTree struct {
	Value int64
	Left  *BTree
	Right *BTree
}

func isSymmetric(node1 *BTree, node2 *BTree) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	return (node1.Value == node2.Value) && isSymmetric(node1.Left, node2.Right) && isSymmetric(node1.Right, node2.Left)
}

func IsSymmetricTree(root *BTree) bool {
	if root == nil {
		return false
	}
	return isSymmetric(root.Left, root.Right)
}
