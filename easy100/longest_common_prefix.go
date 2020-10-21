package easy100

import "math"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minLength := math.MaxInt32
	for _, str := range strs {
		if len(str) < minLength {
			minLength = len(str)
		}
	}
	commonIndex := 0
	for i := 0; i < minLength; i++ {
		item := strs[0][i]
		breakFlag := false
		for j := 1; j < len(strs); j++ {
			if item != strs[j][i] {
				breakFlag = true
				break
			}
		}
		if breakFlag {
			break
		}
		commonIndex++
	}
	if commonIndex > 0 {
		return strs[0][:commonIndex]
	} else {
		return ""
	}
}
