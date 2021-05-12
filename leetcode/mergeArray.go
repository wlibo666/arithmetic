package leetcode

import "fmt"

/*
给定2个排序好的整数数组nums1和nums2，把nums2合并到nums1中成为1个排序的数组。
提示：你可以假定nums1有足够的空间（大小>=m+n）来容纳来自nums2的额外的元素。nums1和nums2的元素的个数各自被初始化为m和n。

分析：
1. 如果nums1从前往后遍历的话，nums2中的元素需要插入nums1，这个时候每插入一次，就会需要将nums1的元素往后移动（或者
需要申请额外的存储空间）。但是我们反过来想，由于合并后的数组长度是确定的，我们可以从最大的数开始写入，这个时候由于
nums1的后面部分的空间是未使用的，刚好可以直接覆写。
2. 需要考虑比较特殊的情况，就是数组可能为空。2个数组都为空自然进不去循环。但是其中一个为空，就要考虑数组可能发生越界的情况了。
*/

func mergeArray(number1 []int, m int, number2 []int, n int) {
	tmpM := m - 1
	tmpN := n - 1
	for p := m + n - 1; p >= 0; p-- {
		fmt.Printf("tmpM:%d,tmpN:%d,p:%d\n", tmpM, tmpN, p)
		if tmpN < 0 || (tmpM >= 0 && number1[tmpM] >= number2[tmpN]) {
			number1[p] = number1[tmpM]
			tmpM--
		} else {
			number1[p] = number2[tmpN]
			tmpN--
		}
	}
}

func MergeArray() {
	n1 := make([]int, 10)
	n2 := make([]int, 4)

	n1[0] = 2
	n1[1] = 3
	n1[2] = 5
	n1[3] = 8

	n2[0] = 1
	n2[1] = 3
	n2[2] = 5
	n2[3] = 7

	fmt.Printf("n1:%v\n", n1)
	fmt.Printf("n2:%v\n", n2)
	mergeArray(n1, 4, n2, 4)
	fmt.Printf("merge n1:%v\n", n1)
}
