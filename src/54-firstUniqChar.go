package main

import "fmt"

//func firstUniqChar(s string) byte {
//	length := len(s)
//	dict := map[byte]int{}
//	for i := 0; i < length; i++{
//		dict[s[i]]++
//	}
//	for i := 0; i < length; i++{
//		if dict[s[i]] == 1{
//			return s[i]
//		}
//	}
//	return ' '
//}

//func firstUniqChar(s string) byte {
//	dict := make([]int, 26)
//	for i := 0; i < len(s); i++ {
//		dict[s[i]-'a']++
//	}
//	for i := 0; i < len(s); i++ {
//		if dict[s[i]-'a'] == 1 {
//			return s[i]
//		}
//	}
//	return ' '
//}

func main() {
	fmt.Println(byte('a'))
}
