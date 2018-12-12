package quick_sort

// cmp 0 1 2 3
// value 3 2 4 5

// 第一次排序,分成两个数组
// 基准值为 3

func QuickSort(data []int) {
	if len(data) < 2 {
		return
	}
	// 第一个数为基准值,从第二个数开始比较，直到head,tail相等,即一遍排序
	base, cmp := data[0], 1
	// 数组起始,结束位置
	head, tail := 0, len(data)-1

	for head < tail {
		if data[cmp] > base {
			// 若大于基准值,则与最后一个数交换
			// 末尾指针前移
			data[cmp], data[tail] = data[tail], data[cmp]
			tail--
		} else {
			// 若小于基准值,则与第一个数交换
			// 起始指针后移
			data[cmp], data[head] = data[head], data[cmp]
			head++
			// 继续比较下一个
			cmp++
		}
	}
	//data[head] = base
	QuickSort(data[:head])
	QuickSort(data[head+1:])
}
