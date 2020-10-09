package medium100

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	result := make([][]int, 0)
	sort.Ints(nums)
	k := len(nums) - 1
	for k >= 2 {
		if nums[k] < 0 {
			break
		}
		target := -1 * nums[k]
		i := 0
		j := k - 1
		for i < j {
			if nums[i]+nums[j] == target {
				elem := make([]int, 0)
				elem = append(elem, nums[i])
				elem = append(elem, nums[j])
				elem = append(elem, target*-1)
				result = append(result, elem)
				for i < j && nums[i] == nums[i+1] {
					i--
				}
				for i < j && nums[j] == nums[j-1] {
					j--
				}
				i++
				j--
			} else if nums[i]+nums[j] < target {
				i++
			} else if nums[i]+nums[j] > target {
				j--
			}
		}
		for k >= 2 && nums[k] == nums[k-1] {
			k--
		}
		k--
	}
	return result
}
