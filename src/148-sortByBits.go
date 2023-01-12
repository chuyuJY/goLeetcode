package main

import "sort"

func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		return cmp(arr[i], arr[j])
	})
	return arr
}

func cmp(num1, num2 int) bool {
	bit1, bit2 := bitCount(num1), bitCount(num2)
	if bit1 < bit2 {
		return true
	} else if bit1 > bit2 {
		return false
	}
	return num1 < num2
}

func bitCount(num int) int {
	count := 0
	for num > 0 {
		num &= (num - 1)
		count++
	}
	return count
}
