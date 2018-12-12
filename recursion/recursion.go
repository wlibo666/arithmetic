package recursion

func Multi(number int) int {
	if number == 1 {
		return number
	}
	// 递归,每次乘以number-1
	return number * Multi(number-1)
}
