package easy100

func romanToInt(s string) int {

	if len(s) == 0 {
		return 0
	}
	var result int
	mapp := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	length := len(s)
	result += mapp[s[length-1]]
	for i := length - 2; i >= 0; i-- {
		cur := s[i]
		pre := s[i+1]
		if mapp[cur] >= mapp[pre] {
			result += mapp[cur]
		} else {
			result -= mapp[cur]
		}
	}
	return result
}
