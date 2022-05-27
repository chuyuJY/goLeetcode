package main

import "fmt"

func sortArray(nums []int) []int {
	temp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		temp[i] = nums[i]
	}
	mergeSort(nums, temp, 0, len(nums)-1)
	return temp
}

func mergeSort(src, dst []int, start, end int) {
	// 一个元素不用
	if start >= end {
		return
	}

	// 1. 划分
	mid := start + (end-start)/2
	// 注意此处dst和src是反过来的，因为合并的时候，是要利用上次的结果
	mergeSort(dst, src, start, mid)
	mergeSort(dst, src, mid+1, end)

	// 2. 归并
	left, right := start, mid+1
	cur := start
	for left <= mid || right <= end {
		if right > end || (left <= mid && src[left] < src[right]) {
			dst[cur] = src[left]
			left++
		} else {
			dst[cur] = src[right]
			right++
		}
		cur++
	}
}

func main() {
	nums := []int{9, 12, 56, 5, 24, 35, 5, 5, 6, 8, 43, 16}
	nums = sortArray(nums)
	fmt.Println(nums)
}
