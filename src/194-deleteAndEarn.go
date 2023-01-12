package main

import "sort"

func deleteAndEarn(nums []int) int {
	// 统计各数字出现频率
	hashMap := map[int]int{}
	for _, num := range nums {
		hashMap[num]++
	}
	// 构建新数组
	newNums := make([]int, 0, len(hashMap))
	for k, _ := range hashMap {
		newNums = append(newNums, k)
	}
	sort.Ints(newNums)

	// 转换为 "打家劫舍"
	dp := make([]int, len(newNums)+1)
	dp[1] = newNums[0] * hashMap[newNums[0]]
	for i := 2; i < len(dp); i++ {
		if newNums[i-2]+1 == newNums[i-1] {
			dp[i] = max(dp[i-2]+newNums[i-1]*hashMap[newNums[i-1]], dp[i-1])
		} else {
			dp[i] = dp[i-1] + newNums[i-1]*hashMap[newNums[i-1]]
		}
	}
	return dp[len(newNums)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
