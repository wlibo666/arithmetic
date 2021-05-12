package leetcode

/*
求浮点数的平方根

解决思路：
 1. 大于等于1的正数n的方根，范围肯定在0~n之间；小于1的正数n的方根，范围肯定在0~1之间
 2. 用二分法（Bisection method, Binary search）从中间开始找n的方根。
 3. 对于大于等于1的正数n，先假设n/2是n的方根，如果n/2的平方大于n，那么说明n的方根在0~n/2之间；如果n/2的平方小于n，说明n的方根在n/2~n之间。以此类推。。
 4. 对于小于1的正数n，先假设0.5是n的方根，方法同上
 这样做的好处是，每次都可以去掉一半可能的值。因此，搜索的范围越来越小。
*/

// 求一个浮点数的平方根，精度为0.0000001
const (
	precision = 0.0000001
)

func abs(number float64) float64 {
	if number >= 0 {
		return number
	}
	return -number
}

func Sqrt(number float64) float64 {
	if number < 0 {
		return -0.1
	}
	tmpNum := number
	for {
		if abs(tmpNum*tmpNum-number) > precision {
			tmpNum = (tmpNum + number/tmpNum) / 2
		} else {
			break
		}
	}
	return tmpNum
}

func HalfSqrt(number float64) float64 {
	var low float64 = 0
	var high float64 = number
	var mid, tmp float64

	for low < high {
		mid = (low + high) / 2
		tmp = mid * mid
		if tmp < (number - precision) {
			low = mid
		} else if tmp > number+precision {
			high = mid
		} else {
			return mid
		}
	}
	return 0.0
}
