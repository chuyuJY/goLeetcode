package main

import "sort"

func lengthOfLIS(nums []int) int {
	dp := []int{}
	for _, num := range nums {
		// sort.SerchInts返回第一个大于num的下标
		if i := sort.SearchInts(dp, num); i == len(dp) {
			dp = append(dp, num)
		} else {
			dp[i] = num
		}
	}
	return len(dp)
}
