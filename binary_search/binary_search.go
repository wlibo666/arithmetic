// Package binarysearch 主要实现二分查找相关算法
package binarysearch

// BinarySearch 二分查找
func BinarySearch(data []int, elem int) (index int) {
	low := 0
	high := len(data) - 1
	mid := 0

	for low <= high {
		mid = (low + high) >> 1
		if data[mid] == elem {
			return mid
		} else if data[mid] > elem {
			// 中间值大，则新的最大值位置为mid-1
			high = mid - 1
		} else {
			// 中间值小，则新的最小值位置为mid+1
			low = mid + 1
		}
	}
	return -1
}
