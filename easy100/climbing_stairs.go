package easy100

func climbStairs(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairsIter(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 1
	}
	result := make([]int, n+1)
	result[0] = 1
	result[1] = 1
	for i := 2; i <= n; i++ {
		result[i] = result[i-1] + result[i-2]
	}
	return result[n]
}
