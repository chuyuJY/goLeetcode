package main

func findMaximumXOR(nums []int) int {
	root := &Trie{}
	root.BuildTree(nums)
	res := 0
	for _, num := range nums {
		// 遍历
		xor := 0
		cur := root
		for i := 31; i >= 0; i-- {
			index := num >> i & 1
			// 存在不同的位，异或为 1
			if cur.next[1-index] != nil {
				xor += 1 << i
				cur = cur.next[1-index]
			} else {
				// 相同
				cur = cur.next[index]
			}
		}
		res = max(res, xor)
	}
	return res
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

type Trie struct {
	next [2]*Trie
}

func (this *Trie) Insert(num int) {
	cur := this
	// 从高位到低位
	for i := 31; i >= 0; i-- {
		index := num >> i & 1
		if cur.next[index] == nil {
			cur.next[index] = &Trie{next: [2]*Trie{}}
		}
		cur = cur.next[index]
	}
}

func (this *Trie) BuildTree(nums []int) {
	for _, num := range nums {
		this.Insert(num)
	}
}
