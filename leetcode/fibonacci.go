package leetcode

import "fmt"

/*
你在爬一个楼梯。到达顶部有n级阶梯。
每次你可以选择爬一级或者二级。在多少不同的方式去到达顶部？

分析：
当n=1，无疑只有一种方式，s=1。
n=2，s=2。
n=3，s=3。
n=4，s=5。
我们发现这个实际是一个斐波那契数列。第一反应是递归实现，f(n)=f(n-1)+f(n-2)，但是递归实现，有多次重复计算，
比如在计算f(n-1)时，需要使用f(n-2)+f(n-3)，但是这个时候另一个递归调用已经去算了f(n-2)，相当于f(n-2)被计算了2次，
这种重复计算的情况普遍存在，将会浪费大量计算时间。很自然的想到反过来操作，递归是从目标算到基础值，
而我们可以采用迭代从基础值1,2累加上去。

*/

// Fibonacci xxx
func Fibonacci(n int) int64 {
	if n == 1 || n == 2 {
		return int64(n)
	}
	return Fibonacci(n-2) + Fibonacci(n-1)
}

func Testfibonacci() {
	fmt.Printf("4!=%d\n", Fibonacci(4))
	fmt.Printf("10!=%d\n", Fibonacci(10))
}
