package easy100

import "math"

func maxProfit(prices []int) int {

	minVal := math.MaxInt32
	maxDiff := 0
	for _, item := range prices {
		if item < minVal {
			minVal = item
		}
		diff := item - minVal
		if diff > maxDiff {
			maxDiff = diff
		}
	}
	return maxDiff
}
