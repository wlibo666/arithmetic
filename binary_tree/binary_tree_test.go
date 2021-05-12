package binary_tree

import "testing"

/*
         1
    2          2
3                   3
*/
func TestIsSymmetricTree(t *testing.T) {
	root := &BTree{Value: 1}
	n1 := &BTree{Value: 2}
	n2 := &BTree{Value: 2}
	n3 := &BTree{Value: 3}
	n4 := &BTree{Value: 3}

	root.Left = n1
	root.Right = n2
	n1.Left = n3
	n2.Right = n4

	t.Logf("isSymmetric：%v", IsSymmetricTree(root))
}
