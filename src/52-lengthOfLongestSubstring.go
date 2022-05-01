package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	l, r, res, length := 0, 0, 1, len(s)
	if length == 0 {
		return 0
	}
	dict := map[byte]int{}
	for r < length {
		if dict[s[r]] == 0 {
			dict[s[r]]++
			r++
			// 只有在r移动的时候，才有可能更大，因此只在此处判断max即可
			res = max(res, r-l)
		} else if dict[s[r]] == 1 {
			dict[s[l]]--
			l++
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func main() {
	str := "pwwkew"
	fmt.Println(lengthOfLongestSubstring(str))
}
