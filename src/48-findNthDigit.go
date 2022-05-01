package main

import (
	"fmt"
	"strconv"
)

func findNthDigit(n int) int {
	if n == 0 {
		return 0
	}
	var start, digit int = 1, 1
	// 1.找到所在位数
	for n > 9*start*digit {
		n -= 9 * start * digit
		start *= 10
		digit += 1
	}
	// 2.找到所在数字    注意此处 -1，因为 n恰好为digit的倍数时，也不会指向下一个（新的）数字
	start += (n - 1) / digit
	// 3.找到该数字的对应位
	// 第一种方法，用降位实现
	//temp := int(math.Pow(10, float64(digit-(n-1+digit)%digit-1)))
	//res := start % (temp * 10) / temp
	// 第二种方法，用转化为字符串，采用下标操作实现
	// 因为go的数组下标不能是负数，因此此处要+digit，，，好吧 字符串下标似乎可以为负，那不用加了。。。
	res := strconv.Itoa(start)[(n-1)%digit] - '0' // 此处-'0'是为了避免ascii码的影响
	//fmt.Printf("%s", res)
	return int(res)
}
func main() {
	// 测试 '0'的ascii码，为48
	//fmt.Printf("%s", '0')
	fmt.Println(findNthDigit(5))
	fmt.Println(10000000 / 2000)
}
