package easy100

func fib(N int) int {
	if N == 0 {
		return 0
	} else if N == 1 {
		return 1
	}
	return fib(N-2) + fib(N-1)
}

func fibIterative(N int) int {
	if N == 0 {
		return 0
	} else if N == 1 {
		return 1
	}
	result := make([]int, N+1)
	result[0] = 0
	result[1] = 1
	for i := 2; i <= N; i++ {
		result[i] = result[i-2] + result[i-1]
	}
	return result[N]
}
