package easy100

func reverseString(s []byte) {
	if len(s) == 0 {
		return
	}
	var tmp byte
	j := len(s) - 1
	for i := 0; i < j; i++ {
		tmp = s[i]
		s[i] = s[j]
		s[j] = tmp
		j--
	}
}
