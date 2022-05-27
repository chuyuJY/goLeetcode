package main

import "fmt"

func singleNonDuplicate(nums []int) int {
	// 偶数下标始终指向一对的第一个元素
	// 按对数进行
	left, right := 0, len(nums)/2
	for left <= right {
		mid := left + (right-left)/2
		// 该对第一个数字的下标
		i := mid * 2
		if i == len(nums)-1 || nums[i] != nums[i+1] {
			if i == 0 || nums[i-2] == nums[i-1] {
				return nums[i]
			}
			// 说明不是第一对配对出错的
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
func main() {
	nums := []int{3, 3, 10, 11, 11}
	fmt.Println(singleNonDuplicate(nums))
}
