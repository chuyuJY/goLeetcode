package main

import "strconv"

func monotoneIncreasingDigits(n int) int {
	// 转成byte切片，方便操作
	nStr := strconv.Itoa(n)
	nBytes := []byte(nStr)
	for i := len(nBytes) - 1; i > 0; i-- {
		// 前一个 若大于 后一个
		if nBytes[i-1] > nBytes[i] {
			nBytes[i-1]--
			for j := i; j < len(nBytes); j++ {
				nBytes[j] = '9'
			}
		}
	}
	res, _ := strconv.Atoi(string(nBytes))
	return res
}
