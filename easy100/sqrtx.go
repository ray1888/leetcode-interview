package easy100

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	start := 1
	end := x
	for start <= end {
		mid := start + (end-start)/2
		if mid*mid > x {
			end = mid - 1
		} else if mid*mid == x {
			return mid
		} else if mid*mid < x {
			start = mid + 1
		}
	}
	return end
}
