package binarysearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	data := []int{1, 2, 3, 5, 7, 8, 9, 10, 23, 30}

	index := BinarySearch(data, 8)
	t.Logf("array:%v\n", data)
	t.Logf("find 8,index:%d\n", index)

	index = BinarySearch(data, 1)
	t.Logf("find 1,index:%d\n", index)

	index = BinarySearch(data, 30)
	t.Logf("find 30,index:%d\n", index)
}
