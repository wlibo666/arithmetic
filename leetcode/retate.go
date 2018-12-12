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
}

func Rotate() {
	numbers := make([][N]int, 0)
	numbers = append(numbers, [N]int{1, 2, 3, 4}, [N]int{5, 6, 7, 8})
	rotate(numbers)
}
