package main

func removeInvalidParentheses(s string) []string {
	hashMap := map[string]bool{}
	path := []byte{}
	// 1. 寻找最少删除的括号数
	left, right := 0, 0
	for _, char := range s {
		if char == '(' {
			left++
		} else if char == ')' {
			if left == 0 {
				right++
			} else {
				left--
			}
		}
	}

	// 2. dfs暴力搜索
	var dfs func(index int, left, right int, leftCnt, rightCnt int)
	dfs = func(index int, left, right int, leftCnt, rightCnt int) {
		if index == len(s) {
			if leftCnt == rightCnt {
				hashMap[string(path)] = true
			}
			return
		}

		// 1. 删去当前元素
		if s[index] == byte('(') {
			// 删去
			if left > 0 {
				dfs(index+1, left-1, right, leftCnt, rightCnt)
			}
		}
		if s[index] == byte(')') {
			if right > 0 {
				dfs(index+1, left, right-1, leftCnt, rightCnt)
			}
		}

		// 2. 保留当前元素
		path = append(path, s[index])
		if s[index] != byte('(') && s[index] != byte(')') {
			dfs(index+1, left, right, leftCnt, rightCnt)
		} else if s[index] == byte('(') {
			dfs(index+1, left, right, leftCnt+1, rightCnt)
		} else if s[index] == byte(')') && leftCnt > rightCnt {
			dfs(index+1, left, right, leftCnt, rightCnt+1)
		}
		path = path[:len(path)-1]
	}

	dfs(0, left, right, 0, 0)
	res := []string{}
	for str, _ := range hashMap {
		res = append(res, str)
	}
	return res
}
