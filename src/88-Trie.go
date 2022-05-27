//package main
//
//import _ "strings"
//
//type Trie struct {
//	next  [26]*Trie // 字母映射表
//	isEnd bool
//}
//
//func Constructor() Trie {
//	return Trie{
//		next:  [26]*Trie{},
//		isEnd: false,
//	}
//}
//
//func (this *Trie) Insert(word string) {
//	cur := this
//	for i := 0; i < len(word); i++ {
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
//func (this *Trie) Search(word string) bool {
//	cur := this
//	for _, char := range word {
//		index := char - 'a'
//		if cur.next[index] == nil {
//			return false
//		}
//		cur = cur.next[index]
//	}
//	return cur.isEnd
//}
//
//func (this *Trie) StartsWith(prefix string) bool {
//	cur := this
//	for i := 0; i < len(prefix); i++ {
//		index := prefix[i] - 'a'
//		if cur.next[index] == nil {
//			return false
//		}
//		cur = cur.next[index]
//	}
//	return true
//}
