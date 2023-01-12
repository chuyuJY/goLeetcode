package main

import (
	"math"
	"sort"
)

func findBestValue(arr []int, target int) int {
	sort.Ints(arr)
	left, right := 0, arr[len(arr)-1]
	for left < right {
		mid := left + (right-left)/2
		if countSum(arr, mid) < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if math.Abs(float64(countSum(arr, left-1)-target)) <= math.Abs(float64(countSum(arr, left)-target)) {
		return left - 1
	}
	return left
}

func countSum(arr []int, base int) int {
	sum := 0
	for _, num := range arr {
		if num > base {
			sum += base
		} else {
			sum += num
		}
	}
	return sum
}
