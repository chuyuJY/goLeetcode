package main

import "fmt"

func findAnagrams(s string, p string) []int {
	res := []int{}
	strTarget := make([]int, 26)
	left, right := 0, 0
	for i := 0; i < len(p); i++ {
		strTarget[p[i]-'a']++
	}
	for right < len(s) {
		temp := &strTarget[s[right]-'a']
		for left < right && *temp == 0 {
			strTarget[s[left]-'a']++
			left++
		}
		if *temp > 0 {
			*temp--
			if right-left+1 == len(p) {
				res = append(res, left)
				strTarget[s[left]-'a']++
				left++
			}
		} else { // 此时，left==right && strTarget[s[right]-'a'] == 0
			left++
		}
		right++
	}
	return res
}
func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))

}
