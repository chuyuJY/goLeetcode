package main

import "fmt"

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	hashMap := make([]int, 'z'-'A'+1)
	tlen := len(t)
	left, right := 0, 0
	for ; right < tlen; right++ {
		hashMap[t[right]-'A']++
		hashMap[s[right]-'A']-- // 第一个窗口
	}
	flag := 0
	for i := 0; i < len(hashMap); i++ {
		if hashMap[i] > 0 {
			flag++
		}
	}
	if flag == 0 {
		return s[:tlen]
	}
	res := s + "#"
	for right <= len(s) {
		for flag == 0 {
			if len(res) > right-left {
				res = s[left:right]
			}
			hashMap[s[left]-'A']++
			if hashMap[s[left]-'A'] > 0 {
				flag++
			}
			left++
		}
		if right < len(s) {
			hashMap[s[right]-'A']--
			if hashMap[s[right]-'A'] == 0 {
				flag--
			}
		}
		right++
	}
	if len(res) > len(s) {
		res = ""
	}
	return res
}
func main() {
	fmt.Println(minWindow("bbaa", "aba"))
}
