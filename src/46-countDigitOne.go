package main

import "fmt"

func countDigitOne(n int) int {
	cnt := 0
	// i是当前位数
	for i := 1; n/i >= 1; i *= 10 {
		// n/i%10 可以得到第i位的数字
		if n/i%10 == 0 {
			cnt += n / (i * 10) * i
		} else if n/i%10 == 1 {
			cnt += n/(i*10)*i + n%i + 1
		} else {
			cnt += n/(i*10)*i + i
		}
	}
	return cnt
}
func main() {
	fmt.Println(countDigitOne(12))
	fmt.Println(countDigitOne(1000))
	fmt.Println(countDigitOne(1))
}
