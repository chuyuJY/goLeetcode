package main

func generateParenthesis(n int) []string {
	res := []string{}
	var backTracking func(left int, right int, combination string)
	backTracking = func(left int, right int, combination string) {
		// 终止条件
		if left == n && right == n {
			res = append(res, combination)
			return
		}
		// 还有左括号未用，可以放左括号
		if left < n {
			backTracking(left+1, right, combination+"(")
		}
		// 已用左括号数量要大于右括号，才能放右括号
		if left > right {
			backTracking(left, right+1, combination+")")
		}
	}
	backTracking(0, 0, "")
	return res
}
