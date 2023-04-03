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
	var n int
	fmt.Scanf("%v\n", &n)
	sc := bufio.NewScanner(os.Stdin)
	nums := []int{}
	if sc.Scan() {
		str := sc.Text()
		strs := strings.Split(str, " ")
		for _, str := range strs {
			num, _ := strconv.Atoi(str)
			nums = append(nums, num)
		}
	}
	colors := ""
	if sc.Scan() {
		colors = sc.Text()
	}
	fmt.Println(test1(nums, colors))
}

//func test1(nums []int, colors string) int64 {
//	var blue, red int64
//	for i := 0; i < len(nums); i++ {
//		if colors[i] == 'B' {
//			blue += int64(nums[i])
//		} else {
//			red += int64(nums[i])
//		}
//	}
//	return (blue % mod) * (red % mod)
//}

func test1(nums []int, colors string) int64 {
	blues, reds := []int64{}, []int64{}
	for i := 0; i < len(nums); i++ {
		if colors[i] == 'B' {
			blues = append(blues, int64(nums[i]))
		} else {
			reds = append(reds, int64(nums[i]))
		}
	}
	var res int64
	for i := 0; i < len(blues); i++ {
		for j := 0; j < len(reds); j++ {
			res += (blues[i] % mod) * (reds[j] % mod)
		}
	}
	return res
}
