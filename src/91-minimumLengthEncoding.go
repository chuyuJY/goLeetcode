//package main
//
//type Trie struct {
//	next  [26]*Trie
//	isEnd bool
//}
//
//func Contructor() Trie {
//	return Trie{
//		next:  [26]*Trie{},
//		isEnd: false,
//	}
//}
//
//func (this *Trie) BuildTrie(words []string) {
//	for _, word := range words {
//		this.Insert(word)
//	}
//}
//
//func (this *Trie) Insert(word string) {
//	cur := this
//	// 反转->构造后缀树
//	for i := len(word) - 1; i >= 0; i-- {
//		index := word[i] - 'a'
//		if cur.next[index] == nil {
//			cur.next[index] = &Trie{
//				next:  [26]*Trie{},
//				isEnd: false,
//			}
//		}
//		cur = cur.next[index]
//	}
//	cur.isEnd = true
//}
//
//func minimumLengthEncoding(words []string) int {
//	root := &Trie{}
//	root.BuildTrie(words)
//	total := 0
//	// 深度优先遍历每一条路径，并统计长度
//	var dfsSumLength func(root *Trie, length int)
//	dfsSumLength = func(root *Trie, length int) {
//		// 寻找到叶子节点
//		isLeaf := true
//		for i := 0; i < 26; i++ {
//			if root.next[i] != nil {
//				isLeaf = false
//				dfsSumLength(root.next[i], length+1)
//			}
//		}
//		if isLeaf {
//			total += length
//		}
//		return
//	}
//	// 起始长度为 1，因为根节点为 #
//	dfsSumLength(root, 1)
//	return total
//}
