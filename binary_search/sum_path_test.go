/*
*    @author wangchunyan.wang
*    @date 1/20/21 6:02 PM
 */

package binarysearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newTree() *Node {
	node11 := &Node{Value: 5}
	node21 := &Node{Value: 4}
	node22 := &Node{Value: 8}
	node31 := &Node{Value: 11}
	node32 := &Node{Value: 13}
	node33 := &Node{Value: 4}
	node41 := &Node{Value: 7}
	node42 := &Node{Value: 2}
	node43 := &Node{Value: 1}

	node11.Left = node21
	node11.Right = node22
	node21.Left = node31
	node22.Left = node32
	node22.Right = node33
	node31.Left = node41
	node31.Right = node42
	node33.Right = node43

	return node11
}

func TestExistPathBySum(t *testing.T) {
	tree := newTree()

	assert.True(t, ExistPathBySum(tree, 22))
	assert.True(t, ExistPathBySum(tree, 18))
	assert.False(t, ExistPathBySum(tree, 20))

	assert.True(t, ExistPathBySum2(tree, 0, 22))
	assert.True(t, ExistPathBySum2(tree, 0, 18))
	assert.False(t, ExistPathBySum2(tree, 0, 20))
}
