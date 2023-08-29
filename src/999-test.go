package main

func main() {
	pivotInteger(8)
}

func pivotInteger(n int) int {
	left, right := 1, n
	for left < right {
		mid := left + (right-left)/2
		if calSum(1, mid) < calSum(mid, n) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if calSum(1, right) != calSum(right, n) {
		return -1
	}
	return right
}

func calSum(start, end int) float64 {
	return float64((start+end)*(end-start+1)) / 2
}
