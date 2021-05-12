package leetcode

import "fmt"

/*
给定一个数组S，它包含n个整数，它是否存在3个元素a，b，c，满足a+b+c=0?找出所有满足条件的元素数组。
提示：a，b，c三个元素必须是升序排列（也就是满足a ≤ b ≤ c），最终的结果不能包含重复的元素数组。例如给定S为{-1 0 1 2 -1 -4}，返回结果是(-1, 0, 1)和(-1, -1, 2)。
For example, given array S = {-1 0 1 2 -1 -4},

A solution set is:
(-1, 0, 1)
(-1, -1, 2)
*/
func ThreeSum(number []int) {
	var res [][]int
	for i := 0; i < len(number); i++ {
		for j := 0; j < len(number); j++ {
			for k := 0; k < len(number); k++ {
				if i != j && i != k && j != k && i <= j && j <= k {
					if number[i]+number[j]+number[k] == 0 {
						arr := make([]int, 0)
						arr = append(arr, number[i])
						arr = append(arr, number[j])
						arr = append(arr, number[k])

						res = append(res, arr)
					}
				}
			}
		}
	}
	if len(res) > 0 {
		for _, arr := range res {
			fmt.Printf("number:%v\n", arr)
		}
	}
}

func ThreeSumTest() {
	numbers := []int{-1, 0, 1, 2, -1, -4}
	ThreeSum(numbers)
}
