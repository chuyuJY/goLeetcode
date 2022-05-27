package main

import (
	"fmt"
	"strings"
)

type Trie struct {
	next  [26]*Trie
	isEnd bool
}

func Constructor() Trie {
	return Trie{
		next:  [26]*Trie{},
		isEnd: false,
	}
}

func (this *Trie) Insert(word string) {
	cur := this
	for _, char := range word {
		index := char - 'a'
		if cur.next[index] == nil {
			cur.next[index] = &Trie{
				next:  [26]*Trie{},
				isEnd: false,
			}
		}
		cur = cur.next[index]
	}
	cur.isEnd = true
}

// 返回前缀
func (this *Trie) FindPrefix(word string) string {
	cur := this
	var flag int
	for ind, char := range word {
		index := char - 'a'
		if cur.next[index] == nil {
			return word
		}
		flag = ind
		cur = cur.next[index]
		// 最短前缀
		if cur.isEnd == true {
			break
		}
	}
	// 返回最短前缀
	return word[:flag+1]
}

func replaceWords(dictionary []string, sentence string) string {
	// 1. 构建前缀树
	prefixTree := Constructor()
	for _, prefix := range dictionary {
		prefixTree.Insert(prefix)
	}
	// 2. 分割句子
	res := strings.Split(sentence, " ")
	for i := 0; i < len(res); i++ {
		res[i] = prefixTree.FindPrefix(res[i])
	}
	// 3. 合并句子
	return strings.Join(res, " ")
}

func main() {
	sentence := []string{"this", "is", "a", "test"}
	fmt.Println(sentence)
}
