package main

// 1. 暴力搜索
// func findWords(board [][]byte, words []string) []string {
// 	res := []string{}
// 	path := []byte{}
// 	hashMap := map[string]bool{}
// 	length := 0
// 	for _, str := range words {
// 		if len(str) > length {
// 			length = len(str)
// 		}
// 		hashMap[str] = true
// 	}
// 	visited := make([][]bool, len(board))
// 	for i := 0; i < len(board); i++ {
// 		visited[i] = make([]bool, len(board[i]))
// 		for j := 0; j < len(board[i]); j++ {
// 			visited[i][j] = false
// 		}
// 	}
// 	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
// 	// dfs暴力搜索
// 	var dfs func(i, j int)
// 	dfs = func(i, j int) {
// 		if hashMap[string(path)] {
// 			res = append(res, string(path))
// 			delete(hashMap, string(path))
// 		}
// 		// 剪枝
// 		if len(path) > length {
// 			return
// 		}

// 		for _, dir := range dirs {
// 			row, col := i+dir[0], j+dir[1]
// 			if 0 <= row && row < len(board) && 0 <= col && col < len(board[0]) && !visited[row][col] {
// 				path = append(path, board[row][col])
// 				visited[row][col] = true
// 				dfs(row, col)
// 				visited[row][col] = false
// 				path = path[:len(path)-1]
// 			}
// 		}
// 	}

// 	// 开始搜索
// 	for i := 0; i < len(board); i++ {
// 		for j := 0; j < len(board[i]); j++ {
// 			path = append(path, board[i][j])
// 			visited[i][j] = true
// 			dfs(i, j)
// 			visited[i][j] = false
// 			path = path[:len(path)-1]
// 		}
// 	}
// 	return res
// }

// 2. 前缀树
type trieNode struct {
	word     string
	children [26]*trieNode
}

func newTrie() *trieNode {
	return &trieNode{}
}

func (tn *trieNode) insert(word string, height int) {
	if len(word) == height {
		tn.word = word
		return
	}
	if tn.children[word[height]-'a'] == nil {
		tn.children[word[height]-'a'] = &trieNode{}
	}
	tn.children[word[height]-'a'].insert(word, height+1)
}

func findWords(board [][]byte, words []string) []string {
	res := []string{}
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[i]))
	}
	trie := newTrie()
	for _, word := range words {
		trie.insert(word, 0)
	}

	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var backTrack func(i, j int, trie *trieNode)
	backTrack = func(i, j int, trie *trieNode) {
		if trie.word != "" {
			res = append(res, trie.word)
			trie.word = "" // 防止重复
		}

		for _, dir := range dirs {
			row, col := i+dir[0], j+dir[1]
			if 0 <= row && row < len(board) && 0 <= col && col < len(board[0]) && !visited[row][col] {
				if trie.children[board[row][col]-'a'] != nil {
					visited[row][col] = true
					backTrack(row, col, trie.children[board[row][col]-'a'])
					visited[row][col] = false
				}
			}
		}
	}

	// 遍历所有起点
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if trie.children[board[i][j]-'a'] != nil {
				visited[i][j] = true
				backTrack(i, j, trie.children[board[i][j]-'a'])
				visited[i][j] = false
			}
		}
	}
	return res
}
