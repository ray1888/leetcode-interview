package easy100

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}
	num1Map := make(map[int]bool)
	for _, item := range nums1 {
		num1Map[item] = true
	}

	resultMap := make(map[int]bool)
	for _, item := range nums2 {
		if _, ok := num1Map[item]; ok {
			resultMap[item] = true
		}
	}
	result := make([]int, 0)
	for item := range resultMap {
		result = append(result, item)
	}
	return result
}
