package main

import (
	"fmt"
)

// 给定一个整形数组和一个整数target，返回2个元素的下标，它们满足相加的和为target。
// 你可以假定每个输入，都会恰好有一个满足条件的返回结果。
/*Example:
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
*/
func TwoSum(numbers []int, target int) (int, int) {
	tmpNumbers := make(map[int]int)
	for index, value := range numbers {
		fmt.Printf("value:%d,will find:%d\n", value, target-value)
		i, ok := tmpNumbers[target-value]
		if ok {
			return i, index
		} else {
			fmt.Printf("set map[%d]%d\n", target-value, index)
			tmpNumbers[value] = index
		}
	}
	return 0, 0
}

func TwoSumTest() {
	number := []int{2, 7, 11, 15}
	i1, i2 := TwoSum(number, 18)
	fmt.Printf("i1:%d,i2:%d\n", i1, i2)
}
