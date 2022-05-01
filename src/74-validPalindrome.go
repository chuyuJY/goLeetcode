package main

import "fmt"

func validPalindrome(s string) bool {
	// 双指针
	left, right := 0, len(s)-1
	isOk := func(l, r int) bool {
		for ; l < r; l, r = l+1, r-1 {
			if s[l] != s[r] {
				return false
			}
		}
		return true
	}
	for ; left < right; left, right = left+1, right-1 {
		if s[left] != s[right] {
			return isOk(left+1, right) || isOk(left, right-1)
		}
	}
	// 若没有不等的情况
	return true
}
func main() {
	fmt.Println(validPalindrome("abccccbca"))
}
