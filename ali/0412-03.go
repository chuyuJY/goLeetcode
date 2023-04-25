package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sc := bufio.NewReader(os.Stdin)
	var str string
	fmt.Fscanln(sc, &str)
	test6(str)
}

func test6(str string) {
	nums := []int{}
	for i := 0; i < len(str); i++ {
		nums = append(nums, int(str[i]-'0'))
	}
	// 求前缀和
	prefixs := []int{}
	cur := 0
	for i := 0; i < len(nums); i++ {
		cur += nums[i]
		prefixs = append(prefixs, cur)
	}
	// 滑动窗口
	res := 0
	sum := prefixs[len(nums)-1]
	for i := 0; i < len(nums); i++ {
		lastNum := nums[len(nums)-1]
		if i != len(nums)-1 && lastNum%5 != 0 {
			continue
		}
		for j := 0; j <= i; j++ {
			if i == len(nums)-1 {
				if j == 0 {
					continue
				} else {
					lastNum = nums[j-1]
				}
			}
			curSum := 0
			if j == 0 {
				curSum = sum - prefixs[i]
			} else {
				curSum = sum - (prefixs[i] - prefixs[j-1])
			}
			// 判断是否有效
			if lastNum%5 == 0 && curSum%3 == 0 {
				res++
			}
		}
	}
	fmt.Println(res)
}
