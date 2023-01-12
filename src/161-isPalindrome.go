package main

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		for left < right && !isWord(s[left]) {
			left++
		}
		for left < right && !isWord(s[right]) {
			right--
		}
		if !isEqual(s[left], s[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

func isWord(char byte) bool {
	return '0' <= char && char <= '9' || 'A' <= char && char <= 'Z' || 'a' <= char && char <= 'z'
}

func isEqual(charA, charB byte) bool {
	return charA|32 == charB|32
}
