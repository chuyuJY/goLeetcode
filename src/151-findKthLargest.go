package main

import "math/rand"

func findKthLargest(nums []int, k int) int {
	fastSort(nums, 0, len(nums)-1, k)
	return nums[len(nums)-k]
}

func fastSort(nums []int, left, right int, k int) {
	if left < right {
		pivot := patition(nums, left, right)
		if pivot == len(nums)-k {
			return
		} else if pivot > len(nums)-k {
			fastSort(nums, left, pivot-1, k)
		} else {
			fastSort(nums, pivot+1, right, k)
		}
	}
}

func patition(nums []int, left, right int) int {
	random := left + rand.Intn(right-left+1)
	nums[random], nums[right] = nums[right], nums[random]
	index := right
	for left < right {
		for left < right && nums[left] < nums[index] {
			left++
		}
		for left < right && nums[right] >= nums[index] {
			right--
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	nums[right], nums[index] = nums[index], nums[right]
	return right
}
