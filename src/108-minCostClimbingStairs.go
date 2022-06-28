package main

func minCostClimbingStairs(cost []int) int {
	// 相当于倒着跳
	pre := cost[1]
	prepre := cost[0]
	for i := 2; i < len(cost); i++ {
		temp := min(pre, prepre) + cost[i]
		pre, prepre = temp, pre
	}
	return min(pre, prepre)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
