package mathx

import "math"

func Max(ints ...int) int {
	max := math.MinInt
	for i := range ints {
		if i > max {
			max = i
		}
	}
	return max
}

func Min(ints ...int) int {
	min := math.MaxInt64
	for i := range ints {
		if i < min {
			min = i
		}
	}
	return min
}
