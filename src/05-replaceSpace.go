package main

import (
	"fmt"
	"regexp"
)

//func replaceSpace(s string) string {
//	newString := strings.Replace(s, " ", "%20", -1)
//	return newString
//}

//func replaceSpace(s string) string {
//	var newString string
//	for _, v := range s {
//		if string(v) == " " {
//			newString += "%20"
//		} else {
//			newString += string(v)
//		}
//	}
//	return newString
//}

//func replaceSpace(s string) string {
//	var newString string
//	for _, v := range s {
//		switch v {
//		case ' ':
//			newString += "%20"
//		default:
//			newString += string(v)
//		}
//	}
//	return newString
//}

// 正则表达式
func replaceSpace(s string) string {
	reg1, _ := regexp.Compile("\\s")
	newString := reg1.ReplaceAllString(s, "%20")
	return newString
}

func main() {
	s := "We are happy."
	fmt.Println(replaceSpace(s))
}
