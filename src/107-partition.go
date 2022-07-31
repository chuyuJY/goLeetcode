package main

import "fmt"

//func partition(s string) [][]string {
//	res := [][]string{}
//	var backTracking func(index int, combination []string)
//	backTracking = func(index int, combination []string) {
//		// 终止条件
//		if index == len(s) {
//			temp := make([]string, len(combination))
//			copy(temp, combination)
//			res = append(res, temp)
//			return
//		}
//		// i为偏移量，也就是要判断[index, index+i]的字符串是否是回文，也即是否可以分割
//		for i := index; i < len(s); i++ {
//			// 剪枝, 若可分割, 那就分割, 并继续往下
//			if isPalindrome(s, index, i) {
//				combination = append(combination, s[index:i+1])
//				backTracking(i+1, combination)
//				// 回溯
//				combination = combination[:len(combination)-1]
//			}
//		}
//	}
//	backTracking(0, []string{})
//	return res
//}
//
//// 判断是否是回文字符串
//func isPalindrome(s string, start, end int) bool {
//	for start < end {
//		if s[start] != s[end] {
//			return false
//		}
//		start++
//		end--
//	}
//	return true
//}

func main() {
	str := "cbbbcc"
	fmt.Println(partition(str))
}
