package main

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	// 定义dp
	dp := make([][]bool, len(s1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(s2)+1)
	}
	// 初始化dp
	// 防止为空的情况
	dp[0][0] = true
	for i := 0; i < len(s1); i++ {
		if s1[i] == s3[i] {
			dp[i+1][0] = dp[i][0]
		}
	}
	for j := 0; j < len(s2); j++ {
		if s2[j] == s3[j] {
			dp[0][j+1] = dp[0][j]
		}
	}
	// 状态转移
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			dp[i+1][j+1] = (s1[i] == s3[i+j+1] && dp[i][j+1]) || (s2[j] == s3[i+j+1] && dp[i+1][j])
		}
	}
	return dp[len(s1)][len(s2)]
}

func main() {
	isInterleave("db", "b", "cbb")
}
