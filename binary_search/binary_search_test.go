package binary_search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	data := []int{1, 2, 3, 5, 7, 8, 9, 10, 23, 30}

	index := BinarySearch(data, 8)
	t.Logf("index:%d", index)
}
