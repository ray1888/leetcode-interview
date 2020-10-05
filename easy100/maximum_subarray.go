package easy100

import "math"

func maxSubArray(nums []int) int {
	maxSum := int(math.Pow(2, 32) * -1)
	sum := 0
	for _, item := range nums {
		if sum > 0 {
			sum += item
		} else {
			sum = item
		}
		if sum > maxSum {
			maxSum = sum
		}
	}
	return maxSum
}
