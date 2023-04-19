package main

import "fmt"

func main() {
	fmt.Println(isValid("123", "123"))
	fmt.Println(isValid("123", "321"))
	fmt.Println(isValid("123", "312"))
}

//func isValid(ori string, dst string) bool {
//	left, right := 0, 0
//	stack := []byte{}
//	for left < len(ori) {
//		if ori[left] == dst[right] {
//			left++
//			right++
//		} else {
//			stack = append(stack, ori[left])
//			left++
//		}
//	}
//	for i := len(stack) - 1; i >= 0; i-- {
//		if stack[i] != dst[right] {
//			return false
//		}
//		right++
//	}
//	return true
//}
