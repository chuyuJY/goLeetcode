package main

import (
	"fmt"
)

//func hammingWeight(num uint32) int {
//	count := 0
//	for i := 0; i < 32; i++ {
//		if 1<<i&num > 0 { // i << i 是把1的二进制向左移动i位，& 是二进制与运算，但返回的是转换后的数字，比如这点是uint，那返回的还是uint
//			count++
//		}
//	}
//	return count
//}

func hammingWeight(num uint32) int {
	count := 0
	for ; num > 0; num >>= 1 { // 和 for ; num > 0; num /= 2 { 等价
		if num&1 == 1 {
			count++
		}
	}
	return count
}

func main() {
	fmt.Printf("%b\n", 9000)
	fmt.Println(hammingWeight(9000))
	//var num uint32 = 500
	//fmt.Println(reflect.TypeOf(num & 12))
}
