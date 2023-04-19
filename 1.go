package main

import (
	"math/rand"
)

/**
 *
 * @param a int整型一维数组
 * @param n int整型
 * @param K int整型
 * @return int整型
 */
func findKth(a []int, n int, K int) int {
	return fastSort(a, 0, n-1, K)
}

func fastSort(nums []int, left, right int, k int) int {
	pivot := partition(nums, left, right)
	if pivot+1 < k {
		return fastSort(nums, pivot+1, right, k)
	} else if pivot+1 > k {
		return fastSort(nums, left, pivot-1, k)
	}
	return nums[k-1]
}

func partition(nums []int, left, right int) int {
	randIndex := left + rand.Intn(right-left+1)
	temp := right
	nums[randIndex], nums[right] = nums[right], nums[randIndex]
	for left < right {
		for left < right && nums[left] < nums[temp] {
			left++
		}
		for left < right && nums[right] >= nums[temp] {
			right--
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	nums[left], nums[temp] = nums[temp], nums[left]
	return left
}
