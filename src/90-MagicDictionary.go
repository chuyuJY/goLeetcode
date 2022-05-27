package main

type MagicDictionary struct {
	next  [26]*MagicDictionary
	isEnd bool
}

func Constructor() MagicDictionary {
	return MagicDictionary{
		next:  [26]*MagicDictionary{},
		isEnd: false,
	}
}

func (this *MagicDictionary) Insert(word string) {
	cur := this
	for _, char := range word {
		index := char - 'a'
		if cur.next[index] == nil {
			cur.next[index] = &MagicDictionary{
				next:  [26]*MagicDictionary{},
				isEnd: false,
			}
		}
		cur = cur.next[index]
	}
	cur.isEnd = true
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, word := range dictionary {
		this.Insert(word)
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	return dfs(this, searchWord, 0, false)
}

// 其中, cur是当前搜索到的下标, flag代表是否已经修改字符
//func dfs(root *MagicDictionary, searchWord string, cur int, flag bool) bool {
//	// 终止条件
//	if root == nil {
//		return false
//	}
//	// 字符串匹配完毕
//	if cur == len(searchWord) {
//		return root.isEnd && flag // 必须修改一次才符合要求
//	}
//	// 还有字符未匹配，遍历所有节点
//	index := searchWord[cur] - 'a'
//	for i := 0; i < 26; i++ {
//		// 节点不为空才进一步判断
//		if root.next[i] != nil {
//			// 1. 若当前节点匹配，不需要修改，继续
//			if int(index) == i {
//				if dfs(root.next[i], searchWord, cur+1, flag) {
//					return true
//				} // 2. 当前节点不匹配且还有修改机会
//			} else if !flag && dfs(root.next[i], searchWord, cur+1, true) {
//				return true
//			}
//		}
//	}
//	return false
//}
