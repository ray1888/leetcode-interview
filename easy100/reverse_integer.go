package easy100

import "math"

func reverse(x int) int {
	var result int

	maxValue := int(math.Pow(2, 31)) - 1
	minValue := -1 * int(math.Pow(2, 31))
	for x != 0 {
		result = (x % 10) + (result * 10)
		x = x / 10
		if result > maxValue || result < minValue {
			return 0
		}
	}
	return result
}
