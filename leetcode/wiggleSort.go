package main

import "fmt"

/*
给你一个没有排序的数组，请将原数组就地重新排列满足如下性质

nums[0] <= nums[1] >= nums[2] <= nums[3]....
Example
Given nums = [3, 5, 2, 1, 6, 4], one possible answer is [1, 6, 2, 5, 3, 4].
*/

func wiggleSort(numbers []int) {
	for i := 1; i < len(numbers); i++ {
		if (i%2 == 1 && numbers[i] < numbers[i-1]) || (i%2 == 0 && numbers[i] > numbers[i-1]) {
			numbers[i-1], numbers[i] = numbers[i], numbers[i-1]
		}
	}
}

func WiggleSort() {
	numbers := []int{3, 5, 2, 1, 6, 4}
	fmt.Printf("array:%v\n", numbers)
	wiggleSort(numbers)
	fmt.Printf("wiggle:%v\n", numbers)
}
