package main

import "fmt"

//func checkInclusion(s1 string, s2 string) bool {
//	strTarget := make([]int, 26)
//	left, right, length := 0, 0, len(s1)
//	for i := 0; i < length; i++ {
//		strTarget[s1[i]-'a']++
//	}
//	for right < len(s2) {
//		cnt := 0
//		for ; right < len(s2) && strTarget[s2[right]-'a'] > 0; right++ {
//			strTarget[s2[right]-'a']--
//			cnt++
//			if cnt == length {
//				return true
//			}
//		}
//		if left != right {
//			strTarget[s2[left]-'a']++
//			left++
//		} else {
//			left++
//			right++
//		}
//	}
//	return false
//}

func checkInclusion(s1 string, s2 string) bool {
	left, right, cnt := 0, 0, 0
	strTarget := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		strTarget[s1[i]-'a']++
	}
	for right < len(s2) {
		// 收缩
		//var temp *int
		temp := &strTarget[s2[right]-'a']
		for ; left <= right && *temp == 0; left++ {
			if cnt > 0 {
				cnt--
				strTarget[s2[left]-'a']++
			}
		}
		if *temp > 0 {
			cnt++
			*temp--
			if cnt == len(s1) {
				return true
			}
		}
		right++
	}
	return false
}
func main() {
	fmt.Println(checkInclusion("hello", "ooolleoooleh"))
}
