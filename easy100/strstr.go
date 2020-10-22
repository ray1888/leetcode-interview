package easy100

func strStr(haystack string, needle string) int {

	if len(needle) == 0 {
		return 0
	} else if len(haystack) == 0 {
		return -1
	}
	result := -1
	i := 0
	for i < (len(haystack) - len(needle) + 1) {
		j := 0
		k := i
		for k < len(haystack) && j < len(needle) {
			if haystack[k] == needle[j] {
				j++
				k++
			} else {
				break
			}
			if j == len(needle) {
				return i
			}
		}
		i++
	}
	return result
}
