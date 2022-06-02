package main

import "fmt"

//func isPalindrome(s string) bool {
//	left, right := 0, len(s)-1
//	for left < right {
//		for ; left < right && !isAlphaNum(s[left]|0b100000); left++ {
//		}
//		for ; left < right && !isAlphaNum(s[right]|0b100000); right-- {
//		}
//		if s[left]|0b100000 != s[right]|0b100000 {
//			return false
//		}
//		left++
//		right--
//	}
//	return true
//}
func isAlphaNum(b byte) bool {
	if b >= 'a' && b <= 'z' || b >= '0' && b <= '9' {
		return true
	}
	return false
}
func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	//fmt.Printf("%v", "0P")
	//str := "A man, a plan, a canal: Panama"
	//fmt.Println("%v", str[23])
}
