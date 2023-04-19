package main

import "fmt"

// [寻找峰值](https://leetcode.cn/problems/find-peak-element/)
// 1. 这是严格递增和递减的情况
func findPeakElement1(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if !(nums[mid] > nums[mid+1]) {
			left = mid + 1
		} else {
			right = mid
		}
	}
	// res := sort.Search(len(nums)-1, func(i int) bool {
	//     return nums[i] > nums[i+1]
	// })
	return right
}

// 2. 这是非严格递增和递减的情况
func findPeakElement2(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] == nums[mid+1] {
			temp := mid
			for temp < right && nums[temp] == nums[temp+1] {
				temp++
			}
			left = temp
			continue
		}
		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[right]
}

func main() {
	fmt.Println(findPeakElement2([]int{2, 3, 2, 2, 2}))
	//fmt.Println(findPeakElement2([]int{2, 2, 3, 2, 2}))
	//fmt.Println(findPeakElement2([]int{2, 2, 2, 3, 2}))
}
