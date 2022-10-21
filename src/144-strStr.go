package main

func strStr(haystack string, needle string) int {
	next := make([]int, len(needle))
	getNext(next, needle)

	// 利用前缀表
	j := 0
	for i := 0; i < len(haystack); i++ {
		for j > 0 && needle[j] != haystack[i] {
			j = next[j-1]
		}
		if needle[j] == haystack[i] {
			j++
		}
		// 若遍历结束模式串
		if j == len(needle) {
			return i - len(needle) + 1
		}
	}
	return -1
}

// 精髓
func getNext(next []int, needle string) {
	j := 0
	next[0] = j
	for i := 1; i < len(needle); i++ {
		for j > 0 && needle[i] != needle[j] {
			j = next[j-1]
		}
		if needle[i] == needle[j] {
			j++
		}
		next[i] = j
	}
}
