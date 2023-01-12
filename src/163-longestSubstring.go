package main

import "strings"

func longestSubstring(s string, k int) int {
	if len(s) < k {
		return 0
	}
	hashMap := map[byte]int{}
	for i := 0; i < len(s); i++ {
		hashMap[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if hashMap[s[i]] < k {
			cnt := 0
			for _, str := range strings.Split(s, string(s[i])) {
				cnt = max(cnt, longestSubstring(str, k))
			}
			return cnt
		}
	}
	return len(s)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
