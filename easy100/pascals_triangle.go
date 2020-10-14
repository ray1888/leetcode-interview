package easy100

func generate(numRows int) [][]int {
	if numRows <= 0 {
		return [][]int{}
	}

	result := make([][]int, numRows)
	result[0] = []int{1}
	for i := 1; i < numRows; i++ {
		tmpResult := make([]int, i+1)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				tmpResult[j] = 1
			} else {
				tmpResult[j] = result[i-1][j-1] + result[i-1][j]
			}
		}
		result[i] = tmpResult
	}
	return result
}
