package main

import "sort"

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%k > 0 {
		return false
	}
	target := sum / k
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	visited := make([]bool, len(nums))
	// index: 当前路径的索引; sum: 当前路径的总和; count: 当前累计路径数
	var backTrack func(index int, sum int, count int) bool
	backTrack = func(index int, sum int, count int) bool {
		if count == k {
			return true
		}
		if sum == target {
			return backTrack(0, 0, count+1)
		}
		for i := index; i < len(nums); i++ {
			// 如果前一个数没用到, 且当前数和前一个数相等, 可以直接剪枝
			if i > 0 && !visited[i-1] && nums[i] == nums[i-1] {
				continue
			}
			if !visited[i] && sum+nums[i] <= target {
				visited[i] = true
				if backTrack(i+1, sum+nums[i], count) {
					return true
				}
				visited[i] = false
			}
		}
		return false
	}
	return backTrack(0, 0, 0)
}
