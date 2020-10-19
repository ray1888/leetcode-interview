package easy100

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	result := -1
	end := len(nums) - 1
	start := 0
	for start <= end {
		// fmt.Println(start, end )
		mid := start + (end-start)/2
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			result = mid
			break
		}
	}
	return result
}
