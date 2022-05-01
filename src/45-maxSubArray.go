package main

import "fmt"

//func maxSubArray(nums []int) int {
//	// max为以当前未知元素结尾的最大值
//	max := nums[0]
//	for i := 1; i < len(nums); i++ {
//		if max < 0 {
//			max = nums[i]
//		} else {
//			max += nums[i]
//			// 更改原数组
//			nums[i] = max
//		}
//	}
//	for i := 0; i < len(nums); i++ {
//		if max < nums[i] {
//			max = nums[i]
//		}
//	}
//	return max
//}
func maxSubArray(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
