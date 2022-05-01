package main

import "fmt"

func exchange(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		for nums[i]%2 != 0 && i < j {
			i++
		}
		for nums[j]%2 != 1 && i < j {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}
func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(exchange(nums))
}
