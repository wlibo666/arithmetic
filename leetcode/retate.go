package main

import "fmt"

const (
	N = 4
)

func show(numbers [][N]int) {
	for _, arr := range numbers {
		for _, n := range arr {
			fmt.Printf("%d    ", n)
		}
		fmt.Printf("\n")
	}
}

/*
二维数组旋转90度
*/
func rotate(numbers [][N]int) {
	show(numbers)
	// 转置
	for i := 0; i < N; i += 1 {
		for j := 0; j < i; j += 1 {
			tmp := numbers[i][j]
			numbers[i][j] = numbers[j][i]
			numbers[j][i] = tmp
		}
	}
	fmt.Printf("after rotate\n")
	show(numbers)

	for i := 0; i < N; i++ {
		for j := N - 1; j >= N/2; j-- {
			tmp := numbers[i][j]
			numbers[i][j] = numbers[i][N-j-1]
			numbers[i][N-j-1] = tmp
		}
	}
	fmt.Printf("after nixu\n")
	show(numbers)
}

func Rotate() {
	numbers := make([][N]int, 0)
	numbers = append(numbers, [N]int{11, 12, 13, 14}, [N]int{15, 16, 17, 18}, [N]int{19, 20, 21, 22}, [N]int{23, 24, 25, 26})
	rotate(numbers)
}
