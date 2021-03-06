package quick_sort

import (
	"testing"
)

func TestQuickSort(t *testing.T) {
	data := []int{6, 8, 3, 6, 10, 9}
	t.Logf("data:%v", data)
	QuickSort(data)
	t.Logf("sort data:%v\n", data)

	data = []int{5, 4, 3, 2, 1}
	t.Logf("data:%v", data)
	QuickSort(data)
	t.Logf("sort data:%v", data)
}
