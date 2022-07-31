package main

import "fmt"

func partition(s string) [][]string {
	res := [][]string{}
	combination := []string{}
	var backTracking func(s string)
	backTracking = func(s string) {
		if len(s) == 0 {
			temp := make([]string, len(combination))
			copy(temp, combination)
			res = append(res, temp)
			return
		}
		for i := 1; i <= len(s); i++ {
			// 剪枝
			prefix := s[:i]
			if isPalindrome(prefix) {
				combination = append(combination, prefix)
				backTracking(s[i:])
				combination = combination[:len(combination)-1]
			}
		}
	}
	backTracking(s)
	return res
}

func isPalindrome(s string) bool {
	start, end := 0, len(s)-1
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func main() {
	str := "aab"
	fmt.Println(partition(str))
}
