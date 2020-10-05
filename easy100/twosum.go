package easy100

func twoSum(nums []int, target int) []int {
	// map[value]index
	digitMap := make(map[int]int)
	for index, number := range nums {
		diff := target - number
		if _, ok := digitMap[diff]; ok {
			return []int{digitMap[diff], index}
		}
		digitMap[number] = index
	}
	return []int{-1}
}
