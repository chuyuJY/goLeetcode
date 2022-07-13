package main

import (
	"math"
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	// 1. 按照绝对值大小，降序排列
	sort.Slice(nums, func(i, j int) bool {
		return math.Abs(float64(nums[i])) > math.Abs(float64(nums[j]))
	})
	// 2. 从最大的负数开始改
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 && k > 0 {
			nums[i] = -nums[i]
			k--
		}
	}
	// 3. 改最小的正数
	if k%2 == 1 {
		nums[len(nums)-1] = -nums[len(nums)-1]
	}
	// 4. 求和
	res := 0
	for i := 0; i < len(nums); i++ {
		res += nums[i]
	}
	return res
}

func main() {
	nums := []int{4, 2, 3}
	largestSumAfterKNegations(nums, 1)
}
