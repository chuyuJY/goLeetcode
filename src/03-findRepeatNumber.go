package main

import "fmt"

// 哈希表
/*func findRepeatNumber(nums []int) int {
	numIndexMap := make(map[int]int)
	for i := range nums {
		index, ok := numIndexMap[nums[i]]
		if ok {
			return nums[index]
		}
		numIndexMap[nums[i]] = i
	}
	return 0
}*/

// 原地置换
func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
			nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
		}
	}
	return -1
}

func main() {
	nums := []int{2, 3, 1, 0, 2, 5, 3}
	fmt.Println(findRepeatNumber(nums))
}
