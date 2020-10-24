package easy100

func singleNumber(nums []int) int {

	result := 0
	for _, item := range nums {
		result = result ^ item
	}
	return result
}
