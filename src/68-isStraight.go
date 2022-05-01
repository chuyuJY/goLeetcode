package main

import (
	"fmt"
	"sort"
)

func isStraight(nums []int) bool {
	sort.Ints(nums)
	i := 0
	for ; i < len(nums) && nums[i] == 0; i++ {
	}
	minNum, maxNum := nums[i], nums[len(nums)-1]
	for ; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			return false
		}
	}
	if maxNum-minNum < 5 {
		return true
	} else {
		return false
	}
}
func main() {
	nums := []int{0, 0, 0, 4, 1}
	fmt.Println(isStraight(nums))
	fmt.Println(nums)
}
