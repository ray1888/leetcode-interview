package easy100

func game(guess []int, answer []int) int {
	count := 0
	for i := 0; i < len(guess); i++ {
		if guess[i] == answer[i] {
			count++
		}
	}
	return count
}
