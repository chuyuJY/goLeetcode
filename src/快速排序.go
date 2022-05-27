package main

import (
	"fmt"
	"math/rand"
)

// 1. 第一种
func fastSort(nums []int, start, end int) []int {
	partition := func(nums []int, start, end int) int {
		// 随机选择一个，并和最右元素交换
		random := start + rand.Intn(end-start+1)
		nums[end], nums[random] = nums[random], nums[end]
		// 存储下标
		index := end
		for start < end {
			for start < end && nums[start] < nums[index] {
				start++
			}
			for start < end && nums[end] >= nums[index] {
				end--
			}
			nums[start], nums[end] = nums[end], nums[start]
		}
		nums[end], nums[index] = nums[index], nums[end]
		return end
	}
	if start < end {
		pivot := partition(nums, start, end)
		fastSort(nums, start, pivot-1)
		fastSort(nums, pivot+1, end)
	}
	return nums
}

// 2. 第二种
//func fastSort(nums []int, left, right int) []int {
//	if left < right {
//		pivot := partition(nums, left, right)
//		fastSort(nums, left, pivot-1)
//		fastSort(nums, pivot+1, right)
//	}
//	return nums
//}
//
//func partition(nums []int, left, right int) int {
//	random := left + rand.Intn(right-left+1)
//	nums[random], nums[right] = nums[right], nums[random]
//
//	small := left - 1
//	for i := left; i < right; i++ {
//		if nums[i] < nums[right] {
//			small++
//			nums[i], nums[small] = nums[small], nums[i]
//		}
//	}
//	small++
//	nums[right], nums[small] = nums[small], nums[right]
//	return small
//}

func main() {
	nums := []int{9, 12, 56, 5, 24, 35, 5, 5, 6, 8, 43, 16}
	nums = fastSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
