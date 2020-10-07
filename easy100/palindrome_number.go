package easy100

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	tmp := x
	y := 0
	for tmp != 0 {
		num := tmp % 10
		y = y*10 + num
		tmp = tmp / 10
	}
	return y == x
}
