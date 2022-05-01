package main

import (
	"fmt"
)

func minArray(numbers []int) int {
	// 定义两个标识
	var (
		low   = 0
		high  = len(numbers) - 1
		pivot int
	)
	for low != high {
		pivot = (low + high) / 2
		if numbers[pivot] > numbers[high] {
			low = pivot + 1
		} else if numbers[pivot] < numbers[high] {
			high = pivot
		} else {
			high--
		}
	}
	return numbers[low]
}

func main() {
	myArray := []int{3, 4, 5, 1, 2}
	fmt.Println(minArray(myArray))
}
