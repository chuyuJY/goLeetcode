package main

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

//func dfs(board [][]byte, word string, i, j, k int) bool {
//	// 先写递归终止条件
//	// 失败的终止条件：越界或者不匹配
//	if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 || board[i][j] != word[k] {
//		return false
//	}
//	// 成功的终止条件：没有失败(通过上述判断)，匹配到最后一个
//	if k == len(word)-1 {
//		return true
//	}
//	// 标记匹配过的字符，防止重复
//	board[i][j] = ' '
//	// 剪枝：往四面八方找就完事了, 因为board[i][j]已经被标记，所以不可能重复匹配
//	res := dfs(board, word, i, j+1, k+1) || dfs(board, word, i+1, j, k+1) || dfs(board, word, i-1, j, k+1) || dfs(board, word, i, j-1, k+1)
//	// 要是没被匹配，还需要还原的
//	if !res {
//		board[i][j] = word[k]
//	}
//	return res
//}
