package main

func maximumJumps(nums []int, target int) int {
	res := -1
	dp := make([]int, len(nums))
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = -1
	}
	for i := 0; i < len(nums); i++ {
		if dp[i] >= 0 {
			for j := i + 1; j < len(nums); j++ {
				if -target <= nums[j]-nums[i] && nums[j]-nums[i] <= target {
					dp[j] = max(dp[j], dp[i]+1)
				}
			}
		}
	}
	res = max(res, dp[len(dp)-1])
	return res
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

//func main() {
//	fmt.Println(maximumJumps([]int{1, 3, 6, 4, 1, 2}, 3))
//}
