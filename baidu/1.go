package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var mod int64 = 1e9 + 7

func main() {
	sc := bufio.NewReader(os.Stdin)
	str, _ := sc.ReadString('\n')
	str = strings.TrimSuffix(str, "\r\n")
	_, _ = strconv.Atoi(str)

	nums := []int64{}
	str, _ = sc.ReadString('\n')
	str = strings.TrimSuffix(str, "\r\n")
	strs := strings.Split(str, " ")
	for _, str := range strs {
		num, _ := strconv.ParseInt(str, 10, 64)
		nums = append(nums, num)
	}
	colors, _ := sc.ReadString('\n')
	str = strings.TrimSuffix(str, "\r\n")
	fmt.Println(test1(nums, colors))
}

func test1(nums []int64, colors string) int64 {
	var blue, red int64
	for i := 0; i < len(nums); i++ {
		if colors[i] == 'B' {
			blue += nums[i]
		} else {
			red += nums[i]
		}
	}
	return ((blue % mod) * (red % mod)) % mod
}

//func test1(nums []int64, colors string) int64 {
//	var res int64
//	blues, reds := []int64{}, []int64{}
//	for i := 0; i < len(nums); i++ {
//		if colors[i] == 'B' {
//			blues = append(blues, nums[i])
//		} else {
//			reds = append(reds, nums[i])
//		}
//	}
//	for i := 0; i < len(blues); i++ {
//		for j := 0; j < len(reds); j++ {
//			res += blues[i] * reds[j] % mod
//		}
//	}
//	return res % mod
//}
