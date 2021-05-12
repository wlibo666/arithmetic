package leetcode

import "testing"

func feibona(n uint32) uint32 {
	if n < 2 {
		return n
	}
	return feibona(n-1) + feibona(n-2)
}

func TestFibonacci(t *testing.T) {
	//TwoSumTest()
	//ThreeSumTest()
	//FibonacciTest()
	//MergeArray()
	//Palindrome()
	//Rotate()
	//WiggleSort()
	//SingleNumber()

	//OrderParseTest()

	/*n := Sqrt(3.25)
	fmt.Printf("3.25 sqrt:%f\n", n)

	n = HalfSqrt(3.25)
	fmt.Printf("3.25 halfsqrt:%f\n", n)*/

	//fmt.Printf("feibona(12):%d\n", feibona(12))

	getDstOrder([]int{1, 12, 2, 8, 3, 11, 4, 9, 5, 13, 6, 10, 7})

	getOriginalOrder([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13})
}
